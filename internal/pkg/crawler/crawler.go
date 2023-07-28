package crawler

import (
	"golang.org/x/net/html"
)

type Crawler interface {
	Crawl(node *html.Node)
}

type crawlerStack[T any] []T

func (cs crawlerStack[T]) push(node T) crawlerStack[T] {
	return append(cs, node)
}

func (cs crawlerStack[T]) pop() (T, crawlerStack[T]) {
	l := len(cs)

	if l == 0 {
		var empty T
		return empty, cs
	}

	return cs[l-1], cs[:l-1]
}

func (cs crawlerStack[T]) len() int {
	return len(cs)
}

// GenericCrawler pre-order traverse HTML crawler.
//
// Calls passed `crawl` handler every time a new HTML ElementNode is found.
func GenericCrawler(node *html.Node, crawl Crawler) {
	if node == nil {
		return
	}

	var top *html.Node
	stack := crawlerStack[*html.Node]{}
	stack = stack.push(node)

	for {
		l := stack.len()

		if l == 0 {
			break
		}

		top, stack = stack.pop()

		if IsElementNode(top) {
			crawl.Crawl(top)
		}

		if e := skipToElementNode(top.NextSibling); e != nil {
			stack = stack.push(e)
		}

		if e := skipToElementNode(top.FirstChild); e != nil {
			stack = stack.push(e)
		}
	}
}

func IsElementNode(node *html.Node) bool {
	return node.Type == html.ElementNode
}

func skipToElementNode(node *html.Node) *html.Node {
	for e := node; e != nil; e = e.NextSibling {
		if IsElementNode(e) {
			return e
		}
	}

	return nil
}
