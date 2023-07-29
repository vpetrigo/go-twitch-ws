package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type eventsubEvent struct {
	Name   string
	Fields []eventsubEventField
}

type eventsubEventField struct {
	FieldName   string
	Name        string
	Type        string
	Description string
}

func newEventsubEventField(name, ty, description string) eventsubEventField {
	titleCase := cases.Title(language.AmericanEnglish)
	splittedName := strings.Split(name, "_")

	for j := range splittedName {
		splittedName[j] = titleCase.String(splittedName[j])
	}

	return eventsubEventField{
		FieldName:   strings.Join(splittedName, ""),
		Name:        name,
		Type:        ty,
		Description: description,
	}
}

func (e *eventsubEvent) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("EventSub Event {%s}\n", e.Name))

	for i, v := range e.Fields {
		sb.WriteString(fmt.Sprintf("  #%2d: [%s] [%s]\n", i+1, v.Name, v.Type))
	}

	return sb.String()
}

func (e *eventsubEvent) stripDescriptionToComment() {
	for i := range e.Fields {
		e.Fields[i].Description = strings.Split(e.Fields[i].Description, ".")[0] + "."
	}
}

func (e *eventsubEvent) updateStructName() {
	e.Name = strings.ReplaceAll(e.Name, " ", "")
}

func (e *eventsubEvent) updateTypeToGoAcceptable() {
	for i := range e.Fields {
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
}
