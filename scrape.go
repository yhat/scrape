package scrape

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Matcher func(node *html.Node) bool

func Find(node *html.Node, matcher Matcher) []*html.Node {
	if matcher(node) {
		return []*html.Node{node}
	}

	matched := []*html.Node{}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		found := Find(c, matcher)
		if len(found) > 0 {
			matched = append(matched, found...)
		}
	}
	return matched
}

func FindOne(node *html.Node, matcher Matcher) (n *html.Node, ok bool) {
	found := Find(node, matcher)
	if len(found) == 1 {
		return found[0], true
	}
	return nil, false
}

func Text(node *html.Node) string {
	joiner := func(s []string) string {
		n := 0
		for i := range s {
			trimmed := strings.TrimSpace(s[i])
			if trimmed != "" {
				s[n] = trimmed
				n++
			}
		}
		return strings.Join(s[:n], " ")
	}
	return TextJoin(node, joiner)
}

func TextJoin(node *html.Node, join func([]string) string) string {
	nodes := Find(node, func(n *html.Node) bool { return n.Type == html.TextNode })
	parts := make([]string, len(nodes))
	for i, n := range nodes {
		parts[i] = n.Data
	}
	return join(parts)
}

func Attr(node *html.Node, key string) string {
	for _, a := range node.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

func ByTag(a atom.Atom) Matcher {
	return func(node *html.Node) bool { return node.DataAtom == a }
}

func ById(id string) Matcher {
	return func(node *html.Node) bool { return Attr(node, "id") == id }
}
