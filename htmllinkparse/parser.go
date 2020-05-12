package htmllinkparse

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func parseHrefs(n *html.Node, links *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		var l Link
		for _, a := range n.Attr {
			if a.Key == "href" {
				//fmt.Println(a.Val)
				l.Href = a.Val
				l.Text = ""

				parseText(n, &l)

				*links = append(*links, l)

				return
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseHrefs(c, links)
	}
}

func parseText(n *html.Node, link *Link){
	if n.Type == html.TextNode {
		link.Text += strings.Join(strings.Fields(n.Data), " ") + " "//strings.TrimSpace(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseText(c, link)
	}
}

func ParseHTML(in io.Reader) ([]Link, error) {
	doc, err := html.Parse(in)
	if err != nil {
		return nil, err
	}

	links := make([]Link, 0)

	parseHrefs(doc, &links)

	return links, nil
}
