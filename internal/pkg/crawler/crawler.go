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

// ElementTraversal pre-order traverse HTML document.
//
// Found HTML ElementNode items passed to `crawler`.
func ElementTraversal(node *html.Node, crawler Crawler) {
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
			crawler.Crawl(top)
		}

		if e := skipToElementNode(top.NextSibling); e != nil {
			stack = stack.push(e)
		}

		if e := skipToElementNode(top.FirstChild); e != nil {
			stack = stack.push(e)
		}
	}
}

// IsElementNode helper for checking the given HTML node is an ElementNode.
func IsElementNode(node *html.Node) bool {
	return node.Type == html.ElementNode
}

// IsTextNode helper for checking the given HTML node is a TextNode.
func IsTextNode(node *html.Node) bool {
	return node.Type == html.TextNode
}

// skipToElementNode helper for skipping to the next ElementNode if available.
func skipToElementNode(node *html.Node) *html.Node {
	for e := node; e != nil; e = e.NextSibling {
		if IsElementNode(e) {
			return e
		}
	}

	return nil
}
