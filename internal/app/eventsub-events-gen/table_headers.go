package main

import (
	"log/slog"

	"golang.org/x/net/html"

	"github.com/vpetrigo/go-twitch-ws/internal/pkg/crawler"
)

var (
	standardTableHeader = []string{
		"Name",
		"Type",
		"Description",
	}
	charityTableHeader = []string{
		"Field",
		"Type",
		"Description",
	}
	conditionTableHeader = []string{
		"Name",
		"Type",
		"Required?",
		"Description",
	}
)

type withTableValidatorFn func(*html.Node) bool

func standardEventTableValidator(tableHeaderNode *html.Node) bool {
	return validateTableHeading(tableHeaderNode, standardTableHeader)
}

func charityEventTableValidator(tableHeaderNode *html.Node) bool {
	return validateTableHeading(tableHeaderNode, charityTableHeader)
}

func conditionTableValidator(tableHeaderNode *html.Node) bool {
	return validateTableHeading(tableHeaderNode, conditionTableHeader)
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

	slog.Debug("validate table heading", "expected", validHeading, "actual", out)

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

func verifyHeader(node *html.Node, checker ...withTableValidatorFn) bool {
	var tr *html.Node

	for tr = node.FirstChild; tr != nil; tr = tr.NextSibling {
		if tr.Data == "tr" {
			break
		}
	}

	if tr == nil {
		slog.Error("nil table row", "node", node)
		return false
	}

	for _, c := range checker {
		if c(tr) {
			return true
		}
	}

	return false
}
