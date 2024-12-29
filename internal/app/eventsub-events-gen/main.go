package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"

	"github.com/vpetrigo/go-twitch-ws/internal/pkg/crawler"
	"github.com/vpetrigo/go-twitch-ws/internal/pkg/refdoc"
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

type eventsubFieldType int

const (
	plainType eventsubFieldType = iota
	compositeType
)

type tableRowProcessor interface {
	RowHandler(tableRow *html.Node)
}

type tableRowProcessWithFn func(*html.Node)

func (t tableRowProcessWithFn) RowHandler(tableRow *html.Node) {
	t(tableRow)
}

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
	_ = generateEventsubFiles(events)
}

func getEvents(resp *html.Node) []eventsubEvent {
	const expectedEventNumber = 80
	events := &eventsubEventCrawler{
		events: make([]eventsubEvent, 0, expectedEventNumber),
	}
	crawler.ElementTraversal(resp, events)
	return events.events
}

func (e *eventsubEventCrawler) Crawl(node *html.Node) {
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
		logrus.Debugf("Found: %s", text.Data)

		if text.Data == "Events" {
			e.state = eventHeaderSearch
			logrus.Debugf("Events Found")
		} else if strings.Contains(text.Data, "Condition") {
			logrus.Debugf("Condition found %s", text.Data)
			processConditions(node)
		}
	}
}

func (e *eventsubEventCrawler) checkEventHeader(node *html.Node) {
	isEventHeader := node.Data == "h3"
	isShoutOutHeader := node.Data == "h2" && strings.Contains(node.FirstChild.Data, "Shoutout")
	isShieldHeader := node.Data == "h2" && strings.Contains(node.FirstChild.Data, "Shield")

	if isEventHeader {
		d := node.FirstChild.Data
		headerText := strings.TrimSuffix(d, "\n")

		if strings.HasSuffix(headerText, "Event") {
			e.tempEvent = newEventsubEvent(headerText)
			e.state = eventTableSearch
			logrus.Tracef("Found: %s", d)
		} else {
			logrus.Errorf("Event ends on %#v", node)
			e.state = eventHeaderSearch
		}
	} else if isShoutOutHeader || isShieldHeader {
		d := node.FirstChild.Data
		headerText := strings.TrimSuffix(d, "\n")
		headerText = fmt.Sprintf("%s Event", headerText)
		e.tempEvent = newEventsubEvent(headerText)
		e.state = eventTableSearch
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
		if !verifyHeader(node, standardEventTableValidator, charityEventTableValidator) {
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
	processor := func(tr *html.Node) {
		field, fieldType := getEventsubFieldFromTable(tr)

		if fieldType == mainField {
			e.tempEvent.addEventField(&field)
		} else {
			e.tempEvent.addInnerEventFieldToLastField(&field)
		}
	}

	tableRawTraverser(node, tableRowProcessWithFn(processor))
	e.events = append(e.events, e.tempEvent)
	e.tempEvent.Fields = nil
	e.state = eventHeaderSearch
}

func getElementValue(node *html.Node) string {
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

	return sb.String()
}

func getFieldTypeRelation(value string) fieldTypeRelation {
	fieldRelation := mainField

	if strings.Contains(value, "\u00a0") {
		fieldRelation = innerField
	}

	return fieldRelation
}

func getFieldTypeDescriptor(node *html.Node) eventsubFieldType {
	for it := node; it != nil; it = it.NextSibling {
		if crawler.IsElementNode(it) && it.Data == "a" {
			for _, a := range it.Attr {
				if a.Key == "href" {
					return compositeType
				}
			}
		}
	}

	return plainType
}

func replaceHTMLSpaces(value string) string {
	return strings.ReplaceAll(value, "\u00a0", "")
}

func getNextPosition(position processPosition) processPosition {
	switch position {
	case namePosition:
		position = typePosition
	case typePosition:
		position = descriptionPosition
	case descriptionPosition:
		position = done
	default:
		panic("unhandled default case")
	}

	return position
}

func generateEventsubFiles(events []eventsubEvent) error {
	const defaultFileAccess = 0o644
	outDir := path.Join("pkg", "eventsub")

	err := writeMainTypes(events, outDir, defaultFileAccess)

	if err != nil {
		return err
	}

	return writeComplexTypes(outDir, defaultFileAccess)
}

func writeComplexTypes(outDir string, defaultFileAccess os.FileMode) error {
	var buf bytes.Buffer
	outFile := path.Join(outDir, "supplementary.go")
	anotherTemplate := path.Join("internal", "app", "eventsub-events-gen", "inner.tmpl")
	t, err := template.ParseFiles(anotherTemplate)

	if err != nil {
		logrus.Error(err)

		return err
	}

	err = t.Execute(&buf, complexTypes)

	if err != nil {
		logrus.Error(err)

		return err
	}

	b, _ := format.Source(buf.Bytes())
	err = os.WriteFile(outFile, b, defaultFileAccess)

	if err != nil {
		logrus.Error(err)

		return err
	}

	return nil
}

func writeMainTypes(events []eventsubEvent, outDir string, defaultFileAccess os.FileMode) error {
	templatePath := path.Join("internal", "app", "eventsub-events-gen", "eventsub.tmpl")
	t, err := template.ParseFiles(templatePath)

	if err != nil {
		logrus.Error(err)

		return err
	}

	for i, e := range events {
		fileName := getFileName(e.SpaceSeparatedName)

		logrus.Tracef("file #%2d: %s", i+1, fileName)
		e.stripDescriptionToComment()
		e.updateTypeToGoAcceptable()

		var buf bytes.Buffer
		err := t.Execute(&buf, e)

		if err != nil {
			logrus.Error(err)

			return err
		}

		b, _ := format.Source(buf.Bytes())
		outFile := path.Join(outDir, fileName)
		_ = os.WriteFile(outFile, b, defaultFileAccess)
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

// getArrayObjectInnerFields handles Drop Entitlement Grant Event that is defined
// in the reference document as several fields in the first table, then <p>...</p> with
// a description and then another table with inner fields description.
func getArrayObjectInnerFields(tr *html.Node) []eventsubEventField {
	var table *html.Node

	for it := tr; it != nil; it = it.Parent {
		if crawler.IsElementNode(it) && it.Data == "table" {
			table = it
			break
		}
	}

	if table == nil {
		panic("table not found")
	}

	// skip to the second table
	for it := table.NextSibling; it != nil; it = it.NextSibling {
		if crawler.IsElementNode(it) && it.Data == "table" {
			table = it
			break
		}
	}

	for it := table.FirstChild; it != nil; it = it.NextSibling {
		if crawler.IsElementNode(it) && it.Data == "tbody" {
			tr = it.FirstChild.NextSibling
			break
		}
	}

	var fields []eventsubEventField

	for it := tr; it != nil; it = it.NextSibling {
		if !crawler.IsElementNode(it) {
			continue
		}

		field, fieldPosition := getEventsubFieldFromTable(it)

		if fieldPosition == mainField {
			fields = append(fields, field)
		} else {
			l := len(fields)
			fields[l-1].InnerFields = append(fields[l-1].InnerFields, field)
		}
	}

	return fields
}

func getEventsubFieldFromTable(tr *html.Node) (eventsubEventField, fieldTypeRelation) {
	var (
		fieldName         string
		fieldTy           string
		fieldDescription  string
		fieldRelation     fieldTypeRelation
		fieldEvent        eventsubEventField
		fieldTyDescriptor eventsubFieldType
	)
	position := namePosition

	for td := tr.FirstChild; position != done && td != nil; td = td.NextSibling {
		if !crawler.IsElementNode(td) {
			continue
		}

		innerTag := td.FirstChild

		if innerTag == nil {
			logrus.Fatalf("Invalid inner tag for %+v", td)
		}

		value := getElementValue(innerTag)
		relation := getFieldTypeRelation(value)
		value = replaceHTMLSpaces(value)

		switch position {
		case namePosition:
			fieldName = strings.TrimSpace(value)
			fieldRelation = relation
		case typePosition:
			fieldTy = strings.ToLower(value)
			fieldTyDescriptor = getFieldTypeDescriptor(innerTag)
		case descriptionPosition:
			fieldDescription = value
		default:
			panic("unhandled default case")
		}

		position = getNextPosition(position)

		if position == done {
			fieldEvent = newEventsubEventField(fieldName, fieldTy, fieldDescription)
			logrus.Tracef("Resulted field: %#v", fieldEvent)
		}
	}

	if fieldTyDescriptor == compositeType {
		fieldEvent.InnerFields = getCompositeType(tr, fieldTy)
	} else if fieldTy == "array" {
		fieldEvent.InnerFields = getArrayObjectInnerFields(tr)
	}

	return fieldEvent, fieldRelation
}

// tableRawTraverser iterates over table's row elements and calls the processor function.
func tableRawTraverser(tableRow *html.Node, processor tableRowProcessor) {
	for tr := tableRow; tr != nil; tr = tr.NextSibling {
		if crawler.IsElementNode(tr) && tr.Data == "tr" {
			processor.RowHandler(tr)
		}
	}
}
