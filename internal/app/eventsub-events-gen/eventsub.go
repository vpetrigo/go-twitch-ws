package main

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type eventsubEvent struct {
	Name               string
	SpaceSeparatedName string
	Fields             []eventsubEventField
}

type eventsubEventField struct {
	FieldName   string
	Name        string
	Type        string
	Description string
	InnerFields []eventsubEventField
}

var complexTypes = map[string]struct {
	Fields []eventsubEventField
}{}

func newEventsubEvent(name string) eventsubEvent {
	return eventsubEvent{
		Name:               strings.ReplaceAll(name, " ", ""),
		SpaceSeparatedName: name,
	}
}

func newEventsubEventField(name, ty, description string) eventsubEventField {
	titleCase := cases.Title(language.AmericanEnglish)
	splitName := strings.Split(name, "_")

	for j := range splitName {
		splitName[j] = titleCase.String(splitName[j])
	}

	fieldName := strings.Join(splitName, "")
	replacePatterns := []struct{ Pattern, Replace string }{
		{"Id", "ID"},
		{"Url", "URL"},
	}

	for _, v := range replacePatterns {
		fieldName = strings.ReplaceAll(fieldName, v.Pattern, v.Replace)
	}

	return eventsubEventField{
		FieldName:   fieldName,
		Name:        name,
		Type:        ty,
		Description: description,
	}
}

func (e *eventsubEvent) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("EventSub Event {%s}\n", e.SpaceSeparatedName))

	for i, v := range e.Fields {
		sb.WriteString(fmt.Sprintf("  #%2d: [%s] [%s]\n", i+1, v.Name, v.Type))
	}

	return sb.String()
}

func (e *eventsubEvent) stripDescriptionToComment() {
	descriptionToComment(e.Fields)
}

func (e *eventsubEvent) updateTypeToGoAcceptable() {
	convertToGoTypes(e.Name, e.Fields)
}

func (e *eventsubEvent) addEventField(field *eventsubEventField) {
	e.Fields = append(e.Fields, *field)
}

func (e *eventsubEvent) addInnerEventFieldToLastField(field *eventsubEventField) {
	l := len(e.Fields)
	e.Fields[l-1].InnerFields = append(e.Fields[l-1].InnerFields, *field)
}

func descriptionToComment(events []eventsubEventField) {
	for i := range events {
		events[i].Description = strings.Split(events[i].Description, ".")[0] + "."

		if len(events[i].InnerFields) > 0 {
			descriptionToComment(events[i].InnerFields)
		}
	}
}

func convertToGoTypes(prefix string, events []eventsubEventField) {
	for i := range events {
		switch events[i].Type {
		case "integer", "int":
			events[i].Type = "int"
		case "boolean":
			events[i].Type = "bool"
		case "string":
		case "string[]":
			events[i].Type = "[]string"
		default:
			if len(events[i].InnerFields) == 0 {
				events[i].Type = "interface{}"
			} else {
				if events[i].Type == "array" {
					events[i].Type = firstLetterToLower(fmt.Sprintf("%s%s", prefix, events[i].FieldName))
				} else {
					events[i].Type = events[i].FieldName
				}

				convertToGoTypes(prefix, events[i].InnerFields)

				if _, ok := complexTypes[events[i].Type]; !ok {
					complexTypes[events[i].Type] = struct{ Fields []eventsubEventField }{events[i].InnerFields}
				}
			}
		}
	}
}

func firstLetterToLower(s string) string {
	if s == "" {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])

	return string(r)
}
