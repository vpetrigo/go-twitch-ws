package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
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

const eventsubTypesFile = "eventsub_types.go"
const eventsubTypesRefURL = "https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/"
const eventsubTypesFileTemplate = `package twitchws

import "github.com/vpetrigo/go-twitch-ws/pkg/eventsub"

type eventSubScope struct {
	Version       string
	MsgType       interface{}
	ConditionType interface{}
}

var (
	eventSubTypes = map[string]eventSubScope{
		{{range .}}"{{.Name}}": {Version: "{{.Version}}", MsgType: &eventsub.{{.MessageType}}{}, ConditionType: &eventsub.{{.ConditionType}}{}},
		{{end}}
	}
)
`

type subscriptionType struct {
	Type        string
	Name        string
	Version     string
	Description string
}

type outputLine struct {
	Name          string
	Version       string
	MessageType   string
	ConditionType string
}

type eventsubCrawler struct {
	Types []subscriptionType
	state eventSubCrawlerState
}

type eventSubCrawlerState uint32

const (
	headingSearch eventSubCrawlerState = iota
	tableSearch
	tableHeadingVerify
	tableBodySearch
	tableRowSearch
	endSearch
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	body, err := refdoc.GetReferenceDocPage(eventsubTypesRefURL)

	if err != nil {
		log.Fatal(err)
	}

	types := getSubscriptionTypes(body)

	if err = generateFile(types); err != nil {
		logrus.Fatal(err)
	}
}

func getSubscriptionTypes(body *html.Node) []subscriptionType {
	types := new(eventsubCrawler)
	crawler.ElementTraversal(body, types)

	return types.Types
}

func (e *eventsubCrawler) Crawl(node *html.Node) {
	if !crawler.IsElementNode(node) || e.state == endSearch {
		return
	}

	switch e.state {
	case headingSearch:
		e.checkSubHeading(node)
	case tableSearch:
		e.checkTableStart(node)
	case tableHeadingVerify:
		e.checkEventsubTableHeading(node)
	case tableBodySearch:
		e.checkTableBodyStart(node)
	case tableRowSearch:
		e.checkRowStart(node)
	}
}

func (e *eventsubCrawler) checkSubHeading(node *html.Node) {
	if node.Data == "h1" {
		text := node.FirstChild

		if text.Data == "Subscription Types" {
			e.state = tableSearch
		}
	}
}

func (e *eventsubCrawler) checkTableStart(node *html.Node) {
	if node.Data == "table" {
		e.state = tableHeadingVerify
	}
}

func (e *eventsubCrawler) checkEventsubTableHeading(node *html.Node) {
	if node.Data == "thead" {
		var tr *html.Node

		for tr = node.FirstChild; tr != nil; tr = tr.NextSibling {
			if tr.Data == "tr" {
				break
			}
		}

		if tr == nil {
			e.state = endSearch
			return
		}

		validHeading := []string{
			"Subscription Type",
			"Name",
			"Version",
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
			e.state = headingSearch
			return
		}

		for i, v := range validHeading {
			if v != out[i] {
				e.state = headingSearch
				return
			}
		}

		e.state = tableBodySearch
	}
}

func (e *eventsubCrawler) checkTableBodyStart(node *html.Node) {
	if node.Data == "tbody" {
		e.state = tableRowSearch
	}
}

func (e *eventsubCrawler) checkRowStart(node *html.Node) {
	validTagsInRow := map[string]struct{}{
		"tr":   {},
		"td":   {},
		"a":    {},
		"code": {},
		"span": {},
		"em":   {},
	}
	const (
		eventsubType = iota
		eventsubName
		eventsubVersion
		eventsubDescription
	)

	if _, ok := validTagsInRow[node.Data]; !ok {
		e.state = endSearch
		return
	}

	evType := subscriptionType{}
	position := eventsubType

	if node.Data == "tr" {
		for td := node.FirstChild; td != nil; td = td.NextSibling {
			if !crawler.IsElementNode(td) {
				continue
			}

			innerTag := td.FirstChild

			if innerTag == nil {
				logrus.Fatalf("Invalid inner tag for %+v", td)
			}

			if innerTag.Data == "span" {
				innerTag = skipToAnchor(innerTag)

				if innerTag == nil {
					log.Fatal("No anchor after span")
				}
			}

			if crawler.IsElementNode(innerTag) {
				tagValue := strings.TrimSuffix(innerTag.Data, "\n")

				switch tagValue {
				case "a", "code":
					value := innerTag.FirstChild.Data

					switch position {
					case eventsubType:
						evType.Type = value
						position = eventsubName
					case eventsubName:
						evType.Name = value
						position = eventsubVersion
					case eventsubVersion:
						evType.Version = value
						position = eventsubDescription
					}

					logrus.Tracef("a/code: %s", value)
				}
			} else {
				if position != eventsubDescription {
					logrus.Fatalf("Invalid position when text found: %d", position)
				}

				var sb strings.Builder

				for text := innerTag; text != nil; text = text.NextSibling {
					if !crawler.IsElementNode(text) {
						sb.WriteString(text.Data)
					} else {
						sb.WriteString(text.FirstChild.Data)
					}
				}

				evType.Description = sb.String()
			}
		}

		logrus.Tracef("Eventsub Type: %#v", evType)
		e.Types = append(e.Types, evType)
	}
}

func skipToAnchor(node *html.Node) *html.Node {
	for it := node.NextSibling; it != nil; it = it.NextSibling {
		if crawler.IsElementNode(it) && it.Data == "a" {
			return it
		}
	}

	return nil
}

func getOutputLines(eventsubTypes []subscriptionType) []outputLine {
	output := make([]outputLine, 0, len(eventsubTypes))

	for _, v := range eventsubTypes {
		splittedName := strings.FieldsFunc(v.Name, func(c rune) bool {
			return c == '.' || c == '_'
		})
		logrus.Trace(splittedName)

		for i := 0; i < len(splittedName); i++ {
			titleCase := cases.Title(language.AmericanEnglish)
			splittedName[i] = titleCase.String(splittedName[i])
		}

		baseName := strings.ReplaceAll(v.Type, " ", "")

		if strings.HasPrefix(baseName, "Goal") {
			baseName = "Goals"
		} else if strings.HasPrefix(baseName, "Shield") {
			baseName = "ShieldMode"
		}

		msgType := fmt.Sprintf("%sEvent", baseName)
		conditionType := fmt.Sprintf("%sEventCondition", baseName)

		output = append(output, outputLine{
			Name:          v.Name,
			Version:       v.Version,
			MessageType:   msgType,
			ConditionType: conditionType,
		})
	}

	return output
}

func generateFile(types []subscriptionType) error {
	const eventsubTypesFilePermissions = 0o644
	outputLines := getOutputLines(types)
	logrus.Debugf("Type size: %d", len(types))
	fileTemplate := template.Must(template.New("eventsub").Parse(eventsubTypesFileTemplate))

	var buf bytes.Buffer
	_ = fileTemplate.Execute(&buf, outputLines)
	b, _ := format.Source(buf.Bytes())

	return os.WriteFile(eventsubTypesFile, b, eventsubTypesFilePermissions)
}
