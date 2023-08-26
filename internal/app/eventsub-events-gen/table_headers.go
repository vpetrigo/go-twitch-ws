package main

import (
	"github.com/sirupsen/logrus"
	"github.com/vpetrigo/go-twitch-ws/internal/pkg/crawler"
	"golang.org/x/net/html"
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
)

func standardEventTableValidator(tableHeaderNode *html.Node) bool {
	return validateTableHeading(tableHeaderNode, standardTableHeader)
}

func charityEventTableValidator(tableHeaderNode *html.Node) bool {
	return validateTableHeading(tableHeaderNode, charityTableHeader)
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
