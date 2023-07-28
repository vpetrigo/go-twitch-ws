package main

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/vpetrigo/go-twitch-ws/internal/pkg/crawler"
	"github.com/vpetrigo/go-twitch-ws/internal/pkg/refdoc"
	"golang.org/x/net/html"
)

const eventsubEventsRefURL = "https://dev.twitch.tv/docs/eventsub/eventsub-reference/"

type eventsubEventField struct {
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

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	resp, err := refdoc.GetReferenceDocPage(eventsubEventsRefURL)

	if err != nil {
		logrus.Fatal(err)
	}

	events := getEvents(resp)
	logrus.Debugf("Events size: %d", len(events))
}

func getEvents(resp *html.Node) []eventsubEvent {
	// TODO: Add support for events with Field/Type/Description table
	const expectedEventNumber = 50
	events := &eventsubEventCrawler{
		events: make([]eventsubEvent, 0, expectedEventNumber),
	}
	crawler.GenericCrawler(resp, events)
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

		validHeading := []string{
			"Name",
			"Type",
			"Description",
		}
		const numberOfHeaderColumns = 4
		out := make([]string, 0, numberOfHeaderColumns)

		for th := tr.FirstChild; th != nil; th = th.NextSibling {
			if !crawler.IsElementNode(th) {
				continue
			}

			out = append(out, th.FirstChild.Data)
		}

		logrus.Tracef("expected: %v, actual: %v", validHeading, out)

		if len(validHeading) != len(out) {
			e.state = eventHeaderSearch
			return
		}

		for i, v := range validHeading {
			if v != out[i] {
				e.state = eventHeaderSearch
				return
			}
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

				if crawler.IsElementNode(innerTag) {
					position = processElementNode(innerTag, &fields, position)
				} else {
					position = processTextNode(innerTag, &fields, position)
				}

				if position == done {
					logrus.Tracef("Resulted field: %#v", fields)
					position = namePosition
					e.tempEvent.Fields = append(e.tempEvent.Fields, fields)
				}
			}
		}
	}

	e.events = append(e.events, e.tempEvent)
	e.state = eventHeaderSearch
}

func processElementNode(node *html.Node, fields *eventsubEventField, position processPosition) processPosition {
	switch position {
	case namePosition:
		tagValue := strings.TrimSuffix(node.Data, "\n")

		if tagValue == "code" {
			fields.Name = node.FirstChild.Data
		}

		position = typePosition
	case typePosition:
		tagValue := strings.TrimSuffix(node.Data, "\n")
		if tagValue == "a" {
			fields.Type = node.FirstChild.Data
		}

		position = descriptionPosition
	default:
		panic(fmt.Errorf("unexpected position: %d", position))
	}

	return position
}

func processTextNode(node *html.Node, fields *eventsubEventField, position processPosition) processPosition {
	switch position {
	case typePosition:
		var sb strings.Builder

		for tag := node; tag != nil; tag = tag.NextSibling {
			if crawler.IsElementNode(tag) {
				sb.WriteString(tag.FirstChild.Data)
			} else if crawler.IsTextNode(tag) {
				sb.WriteString(tag.Data)
			}
		}

		fields.Type = strings.ReplaceAll(sb.String(), "\u00a0", "")
		position = descriptionPosition
	case descriptionPosition:
		var sb strings.Builder

		for tag := node; tag != nil; tag = tag.NextSibling {
			if crawler.IsElementNode(tag) {
				if tag.Data != "br" {
					sb.WriteString(tag.FirstChild.Data)
				} else {
					sb.WriteString("\n")
				}
			} else if crawler.IsTextNode(tag) {
				sb.WriteString(tag.Data)
			}
		}

		fields.Description = sb.String()
		position = done
	default:
		panic(fmt.Errorf("unexpected position: %d", position))
	}

	return position
}
