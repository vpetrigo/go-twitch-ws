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
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const eventsubEventsRefURL = "https://dev.twitch.tv/docs/eventsub/eventsub-reference/"

type eventsubEventField struct {
	FieldName   string
	Name        string
	Type        string
	Description string
}

type crawlerState int
type manyEventsubEventFields []eventsubEventField

type eventsubEvent struct {
	Name   string
	Fields manyEventsubEventFields
}

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
		logrus.Printf("%+v\n", v)
	}

	_ = generateEventsubFiles(events)
}

func (ev eventsubEvent) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("EventSub Event {%s}\n", ev.Name))

	for i, v := range ev.Fields {
		sb.WriteString(fmt.Sprintf("  #%2d: [%s] [%s]\n", i+1, v.Name, v.Type))
	}

	return sb.String()
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
			e.tempEvent.Name = headerText
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
	fields := eventsubEventField{}
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

				position = getElementValue(innerTag, &fields, position)

				if position == done {
					logrus.Tracef("Resulted field: %#v", fields)
					position = namePosition
					e.tempEvent.Fields = append(e.tempEvent.Fields, fields)
				}
			}
		}
	}

	e.events = append(e.events, e.tempEvent)
	e.tempEvent.Fields = nil
	e.state = eventHeaderSearch
}

func getElementValue(node *html.Node, fields *eventsubEventField, position processPosition) processPosition {
	var sb strings.Builder
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

	value := strings.ReplaceAll(sb.String(), "\u00a0", "")

	switch position {
	case namePosition:
		fields.Name = value
		return typePosition
	case typePosition:
		fields.Type = strings.ToLower(value)
		return descriptionPosition
	case descriptionPosition:
		fields.Description = value
		return done
	}

	return position
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

type {{.Name}}Condition struct {}
`

	t := template.Must(template.New("eventsubFileContent").Parse(eventsubFileContentTemplate))

	for i, e := range events {
		splittedEventName := strings.Split(e.Name, " ")
		fileName := getFileName(splittedEventName)

		logrus.Tracef("file #%2d: %s", i+1, fileName)
		e = stripDescriptionToComment(e)
		e = updateStructName(e)
		e = updateTypeToGoAcceptable(e)

		var buf bytes.Buffer
		err := t.Execute(&buf, e)

		if err != nil {
			logrus.Fatal(err)
		}

		b, _ := format.Source(buf.Bytes())
		_, err = os.Stat(fileName)

		if os.IsNotExist(err) {
			_ = os.WriteFile(fileName, b, defaultFileAccess)
		}
	}

	return nil
}

func getFileName(splittedEventName []string) string {
	l := len(splittedEventName)
	loweredEventName := make([]string, 0, l-1)

	for i := 0; i < l-1; i++ {
		loweredEventName = append(loweredEventName, strings.ToLower(splittedEventName[i]))
	}

	return fmt.Sprintf("%s.go", strings.Join(loweredEventName, "_"))
}

func stripDescriptionToComment(e eventsubEvent) eventsubEvent {
	for i := range e.Fields {
		e.Fields[i].Description = strings.Split(e.Fields[i].Description, ".")[0] + "."
	}

	return e
}

func updateStructName(e eventsubEvent) eventsubEvent {
	e.Name = strings.ReplaceAll(e.Name, " ", "")

	return e
}

func updateTypeToGoAcceptable(e eventsubEvent) eventsubEvent {
	for i := range e.Fields {
		titleCase := cases.Title(language.AmericanEnglish)
		splittedName := strings.Split(e.Fields[i].Name, "_")

		for j := range splittedName {
			splittedName[j] = titleCase.String(splittedName[j])
		}

		e.Fields[i].FieldName = strings.Join(splittedName, "")

		switch e.Fields[i].Type {
		case "integer":
			e.Fields[i].Type = "int"
		case "boolean":
			e.Fields[i].Type = "bool"
		case "string":
		case "string[]":
			e.Fields[i].Type = "[]string"
		default:
			e.Fields[i].Type = "interface{}"
		}
	}

	return e
}
