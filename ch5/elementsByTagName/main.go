// Exercise 5.17: a variadic ElementsByTagName that, given an HTML node tree
// and zero or more tag names, returns all elements matching one of the names.
//
//	images   := ElementsByTagName(doc, "img")
//	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
//
// Usage: elementsByTagName <url>
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// ElementsByTagName walks the tree rooted at doc and collects every element
// whose tag name is one of the given names. "name ...string" is variadic, so
// callers may pass any number of tag names.
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node

	// visit is a recursive closure: it checks the current node, then walks
	// each child. Declared first, then assigned, so it can call itself.
	var visit func(n *html.Node)
	visit = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, want := range name {
				if n.Data == want {
					nodes = append(nodes, n)
					break // matched; no need to check the other names
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}

	visit(doc)
	return nodes
}

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "elementsByTagName: %v\n", err)
			continue
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "elementsByTagName: %v\n", err)
			continue
		}

		headings := ElementsByTagName(doc, "h1", "h2", "h3")
		fmt.Printf("%d headings:\n", len(headings))
		for _, h := range headings {
			if h.FirstChild != nil {
				fmt.Printf("  <%s> %s\n", h.Data, h.FirstChild.Data)
			}
		}
	}
}
