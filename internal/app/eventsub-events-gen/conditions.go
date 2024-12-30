package main

import (
	"log/slog"

	"golang.org/x/net/html"

	"github.com/vpetrigo/go-twitch-ws/internal/pkg/crawler"
)

type conditionCrawlerState int

const (
	conditionHeaderSearchState conditionCrawlerState = iota
	conditionFieldsTableSearchState
	conditionFieldsTableVerifyState
	conditionFieldsTableBodySearchState
	conditionFieldsTableParse
	conditionFieldsEndParse
)

type conditionCrawler struct {
	Types map[string]interface{}
	state conditionCrawlerState
}

var isConditionHeadingMet = false

func (c *conditionCrawler) Crawl(node *html.Node) {
	if !crawler.IsElementNode(node) || c.state == conditionFieldsEndParse {
		return
	}

	handlers := map[conditionCrawlerState]func(*conditionCrawler, *html.Node){
		conditionHeaderSearchState:          (*conditionCrawler).headerSearch,
		conditionFieldsTableSearchState:     (*conditionCrawler).tableSearch,
		conditionFieldsTableVerifyState:     (*conditionCrawler).tableVerify,
		conditionFieldsTableBodySearchState: (*conditionCrawler).tableBody,
		conditionFieldsTableParse:           (*conditionCrawler).tableParse,
	}

	if h, ok := handlers[c.state]; ok && h != nil {
		h(c, node)
	}
}

func (c *conditionCrawler) headerSearch(node *html.Node) {
	if node.Data == "h2" {
		if isConditionHeadingMet {
			c.state = conditionFieldsEndParse
			return
		}

		isConditionHeadingMet = true
	}

	if node.Data == "h3" {
		h := node.FirstChild.Data
		c.state = conditionFieldsTableSearchState
		slog.Debug("header", "h", h)
	}
}

func (c *conditionCrawler) tableSearch(node *html.Node) {
	if node.Data == "table" {
		c.state = conditionFieldsTableVerifyState
		slog.Debug("condition table found")
	}
}

func (c *conditionCrawler) tableVerify(node *html.Node) {
	if node.Data == "thead" {
		if !verifyHeader(node, conditionTableValidator) {
			c.state = conditionFieldsEndParse
			return
		}

		c.state = conditionFieldsTableBodySearchState
	}
}

func (c *conditionCrawler) tableBody(node *html.Node) {
	if node.Data == "tbody" {
		c.state = conditionFieldsTableParse
	}
}

func (c *conditionCrawler) tableParse(node *html.Node) {
	processor := func(tableRow *html.Node) {
		field, fieldType := getEventsubFieldFromTable(tableRow)
		slog.Debug("parse table", "field", field, "field type", fieldType)
		if fieldType != mainField {
			panic("unexpected inner field in the condition type")
		}
	}

	tableRawTraverser(node, tableRowProcessWithFn(processor))
	c.state = conditionHeaderSearchState
}

func processConditions(node *html.Node) {
	conditions := &conditionCrawler{
		Types: make(map[string]interface{}),
	}

	crawler.ElementTraversal(node, conditions)
}
