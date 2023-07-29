package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"strings"
	"text/template"

	"github.com/sirupsen/logrus"
	"github.com/vpetrigo/go-twitch-ws/internal/pkg/crawler"
	"github.com/vpetrigo/go-twitch-ws/internal/pkg/refdoc"
	"golang.org/x/net/html"
)

const eventsubEventsRefURL = "https://dev.twitch.tv/docs/eventsub/eventsub-reference/"

type crawlerState int

type eventsubEventCrawler struct {
	events    []eventsubEvent
	tempEvent eventsubEvent
	state     crawlerState
}

const (
	headingSearch crawlerState = iota
	eventHeaderSearch
	eventTableSearch
	eventTableVerify
	eventTableBodySearch
	eventTableParse
	endSearch
)

type processPosition int

const (
	namePosition processPosition = iota
	typePosition
	descriptionPosition
	done
)

type fieldTypeRelation int

const (
	mainField fieldTypeRelation = iota
	innerField
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	resp, err := refdoc.GetReferenceDocPage(eventsubEventsRefURL)

	if err != nil {
		logrus.Fatal(err)
	}

	events := getEvents(resp)
	logrus.Debugf("Events size: %d", len(events))

	for i, v := range events {
		logrus.Printf("Event #%d\n", i+1)
		logrus.Printf("%s\n", &v)
	}

	_ = generateEventsubFiles(events)
}

func getEvents(resp *html.Node) []eventsubEvent {
	const expectedEventNumber = 50
	events := &eventsubEventCrawler{
		events: make([]eventsubEvent, 0, expectedEventNumber),
	}
	crawler.ElementTraversal(resp, events)
	return events.events
}

func (e *eventsubEventCrawler) Crawl(node *html.Node) {
	if crawler.IsElementNode(node) && e.state != endSearch {
		if node.Data == "h3" {
			text := node.FirstChild.Data
			logrus.Tracef("Event found: %s - state: %d", text, e.state)
		}
	}

	if !crawler.IsElementNode(node) || e.state == endSearch {
		return
	}

	handlers := map[crawlerState]func(*eventsubEventCrawler, *html.Node){
		headingSearch:        (*eventsubEventCrawler).checkMainEventHeader,
		eventHeaderSearch:    (*eventsubEventCrawler).checkEventHeader,
		eventTableSearch:     (*eventsubEventCrawler).checkEventTable,
		eventTableVerify:     (*eventsubEventCrawler).verifyEventTable,
		eventTableBodySearch: (*eventsubEventCrawler).checkTableBodyStart,
		eventTableParse:      (*eventsubEventCrawler).parseEventTable,
	}

	if h, ok := handlers[e.state]; ok {
		h(e, node)
	}
}

func (e *eventsubEventCrawler) checkMainEventHeader(node *html.Node) {
	if node.Data == "h2" {
		text := node.FirstChild

		if text.Data == "Events" {
			e.state = eventHeaderSearch
			logrus.Trace("Events Found")
		}
	}
}

func (e *eventsubEventCrawler) checkEventHeader(node *html.Node) {
	if node.Data == "h3" {
		headerText := strings.TrimSuffix(node.FirstChild.Data, "\n")

		if strings.HasSuffix(headerText, "Event") {
			e.tempEvent = newEventsubEvent(headerText)
			e.state = eventTableSearch
			logrus.Tracef("Found: %s", node.FirstChild.Data)
		} else {
			logrus.Errorf("Event ends on %#v", node)
			e.state = endSearch
		}
	}
}

func (e *eventsubEventCrawler) checkEventTable(node *html.Node) {
	if node.Data == "table" {
		e.state = eventTableVerify
		logrus.Trace("Table Found")
	}
}

func (e *eventsubEventCrawler) verifyEventTable(node *html.Node) {
	if node.Data == "thead" {
		var tr *html.Node

		for tr = node.FirstChild; tr != nil; tr = tr.NextSibling {
			if tr.Data == "tr" {
				break
			}
		}

		if tr == nil {
			logrus.Errorf("nil table row: %#v", node)
			e.state = endSearch
			return
		}

		if !(standardEventTableValidator(tr) || charityEventTableValidator(tr)) {
			e.state = eventHeaderSearch
			return
		}

		e.state = eventTableBodySearch
	}
}

func (e *eventsubEventCrawler) checkTableBodyStart(node *html.Node) {
	if node.Data == "tbody" {
		e.state = eventTableParse
	}
}

func (e *eventsubEventCrawler) parseEventTable(node *html.Node) {
	var (
		fieldName        string
		fieldTy          string
		fieldDescription string
		fieldRelation    fieldTypeRelation
	)
	position := namePosition

	for tr := node; tr != nil; tr = tr.NextSibling {
		if crawler.IsElementNode(tr) && tr.Data == "tr" {
			for td := tr.FirstChild; td != nil; td = td.NextSibling {
				if !crawler.IsElementNode(td) {
					continue
				}

				innerTag := td.FirstChild

				if innerTag == nil {
					logrus.Fatalf("Invalid inner tag for %+v", td)
				}

				nextPosition, retValue, relation := getElementValue(innerTag, position)

				switch position {
				case namePosition:
					fieldName = retValue
					fieldRelation = relation
					logrus.Debugf("Relation: [%s] %s %d", e.tempEvent.SpaceSeparatedName, fieldName, relation)
				case typePosition:
					fieldTy = retValue
				case descriptionPosition:
					fieldDescription = retValue
				}

				position = nextPosition

				if position == done {
					field := newEventsubEventField(fieldName, fieldTy, fieldDescription)
					logrus.Tracef("Resulted field: %#v", field)
					position = namePosition

					if fieldRelation == mainField {
						e.tempEvent.Fields = append(e.tempEvent.Fields, field)
					} else {
						l := len(e.tempEvent.Fields)
						e.tempEvent.Fields[l-1].Fields = append(e.tempEvent.Fields[l-1].Fields, field)
					}
				}
			}
		}
	}

	e.events = append(e.events, e.tempEvent)
	e.tempEvent.Fields = nil
	e.state = eventHeaderSearch
}

// TODO: split this function into multiple based on required feature:
// - extract column value
// - determine main/inner field
// - remove HTML &nbsp; characters
// getElementValue extract field element value and specify which position should be updated next.
func getElementValue(node *html.Node, position processPosition) (processPosition, string, fieldTypeRelation) {
	var sb strings.Builder
	fieldRelation := mainField

	for tag := node; tag != nil; tag = tag.NextSibling {
		var value string

		if crawler.IsElementNode(tag) {
			if tag.Data != "br" {
				value = strings.TrimSuffix(tag.FirstChild.Data, "\n")
			} else {
				value = "\n"
			}
		} else if crawler.IsTextNode(tag) {
			value = strings.TrimSuffix(tag.Data, "\n")
		}

		sb.WriteString(value)
	}

	if strings.Contains(sb.String(), "\u00a0") {
		fieldRelation = innerField
	}

	value := strings.ReplaceAll(sb.String(), "\u00a0", "")
	var ret string

	switch position {
	case namePosition:
		ret = value
		position = typePosition
	case typePosition:
		ret = strings.ToLower(value)
		position = descriptionPosition
	case descriptionPosition:
		ret = value
		position = done
	}

	return position, ret, fieldRelation
}

func standardEventTableValidator(tableHeaderNode *html.Node) bool {
	validHeading := []string{
		"Name",
		"Type",
		"Description",
	}

	return validateTableHeading(tableHeaderNode, validHeading)
}

func charityEventTableValidator(tableHeaderNode *html.Node) bool {
	validHeading := []string{
		"Field",
		"Type",
		"Description",
	}

	return validateTableHeading(tableHeaderNode, validHeading)
}

func validateTableHeading(tr *html.Node, validHeading []string) bool {
	validationSliceLen := len(validHeading)
	out := make([]string, 0, validationSliceLen)

	for th := tr.FirstChild; th != nil; th = th.NextSibling {
		if !crawler.IsElementNode(th) {
			continue
		}

		out = append(out, th.FirstChild.Data)
	}

	logrus.Tracef("expected: %v, actual: %v", validHeading, out)

	if validationSliceLen != len(out) {
		return false
	}

	for i, v := range validHeading {
		if v != out[i] {
			return false
		}
	}

	return true
}

func generateEventsubFiles(events []eventsubEvent) error {
	const defaultFileAccess = 0o644
	const eventsubFileContentTemplate = `package twitchws

type {{.Name}} struct {
{{range .Fields}}{{.FieldName}} {{.Type}} ` + "`json:\"{{.Name}}\"` // {{.Description}}\n" + `{{end}}}

{{range .Fields}}{{if .Fields}}type {{.Type}} struct {
{{range .Fields}}{{.FieldName}} {{.Type}} ` + "`json:\"{{.Name}}\"` // {{.Description}}\n" + `{{end}}}

{{end}}{{end}}

type {{.Name}}Condition struct {}
`

	t := template.Must(template.New("eventsubFileContent").Parse(eventsubFileContentTemplate))

	for i, e := range events {
		fileName := getFileName(e.SpaceSeparatedName)

		logrus.Tracef("file #%2d: %s", i+1, fileName)
		e.stripDescriptionToComment()
		e.updateTypeToGoAcceptable()

		var buf bytes.Buffer
		err := t.Execute(&buf, e)

		if err != nil {
			logrus.Fatal(err)
		}

		if strings.HasPrefix(e.Name, "Charity") {
			logrus.Debugf("%s", buf.String())
		}

		b, _ := format.Source(buf.Bytes())
		_, err = os.Stat(fileName)

		if os.IsNotExist(err) {
			_ = os.WriteFile(fileName, b, defaultFileAccess)
		}
	}

	return nil
}

func getFileName(name string) string {
	splitEventName := strings.Split(name, " ")
	l := len(splitEventName)
	loweredEventName := make([]string, 0, l-1)

	for i := 0; i < l-1; i++ {
		loweredEventName = append(loweredEventName, strings.ToLower(splitEventName[i]))
	}

	return fmt.Sprintf("%s.go", strings.Join(loweredEventName, "_"))
}
