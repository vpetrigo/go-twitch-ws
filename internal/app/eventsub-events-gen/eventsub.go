package main

import (
	"fmt"
	"slices"
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
	Type        string // Type to be used as a struct name (if required)
	InnerType   string // InnerType represents a type name to be used inside a struct as a field type
	Description string
	InnerFields []eventsubEventField // Inner fields (if any) for complex types
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
	fieldName := convertUnderscoreSeparated(name)
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
			events[i].InnerType = "int"
		case "boolean", "bool":
			events[i].InnerType = "bool"
		case "string":
			events[i].InnerType = "string"
		case "string[]":
			events[i].InnerType = "[]string"
		default:
			if len(events[i].InnerFields) == 0 {
				events[i].InnerType = "interface{}"
			} else {
				convertComplexTypesToGoTypes(prefix, &events[i])
			}
		}
	}
}

func convertComplexTypesToGoTypes(prefix string, event *eventsubEventField) {
	if len(event.InnerFields) == 0 {
		event.InnerType = "interface{}"
	} else {
		switch event.Type {
		case "array":
			tyName := firstLetterToLower(fmt.Sprintf("%s%s", prefix, event.FieldName))
			event.InnerType = tyName
			event.Type = tyName
		case "object":
			event.InnerType = event.FieldName
			event.Type = event.FieldName
		default:
			arrayTypes := []string{"top_contributions"}
			tyName := convertUnderscoreSeparated(event.Type)
			event.Type = tyName
			event.InnerType = tyName

			if strings.Contains(event.Description, "array") || slices.Contains(arrayTypes, event.Name) {
				event.InnerType = fmt.Sprintf("[]%s", event.InnerType)
			}
		}

		convertToGoTypes(prefix, event.InnerFields)

		if _, ok := complexTypes[event.Type]; !ok {
			complexTypes[event.Type] = struct{ Fields []eventsubEventField }{
				event.InnerFields,
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

func convertUnderscoreSeparated(input string) string {
	titleCase := cases.Title(language.AmericanEnglish)
	split := strings.Split(input, "_")

	for j := range split {
		split[j] = titleCase.String(split[j])
	}

	return strings.Join(split, "")
}
