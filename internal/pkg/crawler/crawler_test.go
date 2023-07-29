package crawler

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type testCrawler func(node *html.Node)

func TestElementTraversal(t *testing.T) {
	const htm = `<!DOCTYPE html>
<html>
<head>
	<title></title>
</head>
<body>
	body content
	<p>more content</p>
</body>
</html>`
	expectedOrder := []string{"html", "head", "title", "body", "p"}
	actualOrder := make([]string, 0, 5)
	doc, _ := html.Parse(strings.NewReader(htm))

	ElementTraversal(doc, newTestCrawler(func(node *html.Node) {
		if IsElementNode(node) {
			actualOrder = append(actualOrder, node.Data)
		}
	}))

	if len(expectedOrder) != len(actualOrder) {
		t.Fatalf("Incorrect length after traversal: expected %d, actual %d", len(expectedOrder), len(actualOrder))
	}

	for i, v := range actualOrder {
		if v != expectedOrder[i] {
			t.Fatalf("Values mismatch on position [%d]: expected %v, actual %v", i, expectedOrder, actualOrder)
		}
	}
}

func TestIsElementNode(t *testing.T) {
	fixture := []struct {
		element  *html.Node
		expected bool
	}{
		{
			element:  createDummyHTMLNode(html.ErrorNode),
			expected: false,
		},
		{
			element:  createDummyHTMLNode(html.TextNode),
			expected: false,
		},
		{
			element:  createDummyHTMLNode(html.DocumentNode),
			expected: false,
		},
		{
			element:  createDummyHTMLNode(html.ElementNode),
			expected: true,
		},
		{
			element:  createDummyHTMLNode(html.CommentNode),
			expected: false,
		},
		{
			element:  createDummyHTMLNode(html.DoctypeNode),
			expected: false,
		},
		{
			element:  createDummyHTMLNode(html.RawNode),
			expected: false,
		},
	}

	for _, v := range fixture {
		result := IsElementNode(v.element)

		if v.expected != result {
			t.Fatalf("IsElementNode(%v) - expected: %t, actual: %t", v.element, v.expected, result)
		}
	}
}

func TestIsTextNode(t *testing.T) {
	fixture := []struct {
		element  *html.Node
		expected bool
	}{
		{
			element:  createDummyHTMLNode(html.ErrorNode),
			expected: false,
		},
		{
			element:  createDummyHTMLNode(html.TextNode),
			expected: true,
		},
		{
			element:  createDummyHTMLNode(html.DocumentNode),
			expected: false,
		},
		{
			element:  createDummyHTMLNode(html.ElementNode),
			expected: false,
		},
		{
			element:  createDummyHTMLNode(html.CommentNode),
			expected: false,
		},
		{
			element:  createDummyHTMLNode(html.DoctypeNode),
			expected: false,
		},
		{
			element:  createDummyHTMLNode(html.RawNode),
			expected: false,
		},
	}

	for _, v := range fixture {
		result := IsTextNode(v.element)

		if v.expected != result {
			t.Fatalf("IsTextNode(%+v) - expected: %t, actual: %t", v.element, v.expected, result)
		}
	}
}

func createDummyHTMLNode(ty html.NodeType) *html.Node {
	return &html.Node{
		Parent:      nil,
		FirstChild:  nil,
		LastChild:   nil,
		PrevSibling: nil,
		NextSibling: nil,
		Type:        ty,
		DataAtom:    atom.A,
		Data:        "",
		Namespace:   "",
	}
}

func (f testCrawler) Crawl(node *html.Node) {
	f(node)
}

func newTestCrawler(fn testCrawler) testCrawler {
	return fn
}
