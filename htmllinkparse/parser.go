package htmllinkparse

import (
	"io"
	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseHTML(in io.Reader) ([]Link, error) {
	doc, err := html.Parse(in)
	if err != nil {
		return nil, err
	}

	links := make([]Link, 0)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					//fmt.Println(a.Val)
					links = append(links, Link{a.Val, "WIP"})
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return links, nil
}
