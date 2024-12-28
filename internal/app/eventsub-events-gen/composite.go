package main

import (
	"strings"

	"golang.org/x/net/html"

	"github.com/vpetrigo/go-twitch-ws/internal/pkg/crawler"
)

type compositeTypeParseState int

const (
	compositeHeaderSearch compositeTypeParseState = iota
	compositeTableSearch
	compositeTableVerify
	compositeTableBodySearch
	compositeParse
	compositeEndSearch
)

type compositeTypeCrawler struct {
	Fields []eventsubEventField
	tyName string
	state  compositeTypeParseState
}

func newCompositeTypeCrawler(tyName string) *compositeTypeCrawler {
	return &compositeTypeCrawler{
		tyName: tyName,
		state:  compositeHeaderSearch,
	}
}

func (c *compositeTypeCrawler) Crawl(node *html.Node) {
	if !crawler.IsElementNode(node) || c.state == compositeEndSearch {
		return
	}

	handlers := map[compositeTypeParseState]func(*compositeTypeCrawler, *html.Node){
		compositeHeaderSearch:    (*compositeTypeCrawler).compositeTypeHeaderHandler,
		compositeTableSearch:     (*compositeTypeCrawler).compositeTypeTableSearchHandler,
		compositeTableVerify:     (*compositeTypeCrawler).compositeTypeTableVerifyHandler,
		compositeTableBodySearch: (*compositeTypeCrawler).compositeTypeTableBodySearchHandler,
		compositeParse:           (*compositeTypeCrawler).compositeTypeTableParseHandler,
	}

	if h, ok := handlers[c.state]; ok {
		h(c, node)
	}
}

func getCompositeType(node *html.Node, tyName string) []eventsubEventField {
	top := returnToParent(node)
	c := newCompositeTypeCrawler(tyName)

	crawler.ElementTraversal(top, c)

	return c.Fields
}

func returnToParent(node *html.Node) *html.Node {
	it := node

	for {
		if it.Parent == nil {
			return it
		}

		it = it.Parent
	}
}

func (c *compositeTypeCrawler) compositeTypeHeaderHandler(node *html.Node) {
	if node.Data == "h2" {
		typeID := strings.ReplaceAll(node.Attr[0].Val, "-", "_")

		if c.tyName == typeID {
			c.state = compositeTableSearch
		}
	}
}

func (c *compositeTypeCrawler) compositeTypeTableSearchHandler(node *html.Node) {
	if node.Data == "table" {
		c.state = compositeTableVerify
	}
}

func (c *compositeTypeCrawler) compositeTypeTableVerifyHandler(node *html.Node) {
	if node.Data == "thead" {
		if !verifyHeader(node, standardEventTableValidator) {
			c.state = compositeEndSearch
			return
		}

		c.state = compositeTableBodySearch
	}
}

func (c *compositeTypeCrawler) compositeTypeTableBodySearchHandler(node *html.Node) {
	if node.Data == "tbody" {
		c.state = compositeParse
	}
}

func (c *compositeTypeCrawler) compositeTypeTableParseHandler(node *html.Node) {
	processor := func(tableRow *html.Node) {
		field, fieldType := getEventsubFieldFromTable(tableRow)

		if fieldType == mainField {
			c.Fields = append(c.Fields, field)
		} else {
			panic("unexpected inner field in the composite type")
		}
	}

	tableRawTraverser(node, tableRowProcessWithFn(processor))
	c.state = compositeEndSearch
}
