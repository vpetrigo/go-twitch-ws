package main

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/vpetrigo/go-twitch-ws/internal/pkg/crawler"
	"golang.org/x/net/html"
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
		var tr *html.Node

		for tr = node.FirstChild; tr != nil; tr = tr.NextSibling {
			if tr.Data == "tr" {
				break
			}
		}

		if tr == nil {
			logrus.Errorf("nil table row: %#v", node)
			c.state = compositeEndSearch
			return
		}

		if !standardEventTableValidator(tr) {
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
	for tr := node; tr != nil; tr = tr.NextSibling {
		if crawler.IsElementNode(tr) && tr.Data == "tr" {
			field, fieldType := getEventsubFieldFromTable(tr)

			if fieldType == mainField {
				c.Fields = append(c.Fields, field)
			} else {
				panic("unexpected inner field in the composite type")
			}
		}
	}

	c.state = compositeEndSearch
}
