package sitemap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func parseHrefs(n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				//check duplicates and host validity

				//no need to investigate children
				return []string{a.Val}
			}
		}
	}
	var s []string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s = append(s, parseHrefs(c)...)
	}

	return s
}

func ParseHTML(in io.Reader) ([]string, error) {
	doc, err := html.Parse(in)
	if err != nil {
		return nil, err
	}

	links := parseHrefs(doc)

	return links, nil
}

func InspectURL(u url.URL, visited map[string]struct{}, depth uint, maxDepth uint) error {
	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = resp.Body.Close()
	if err != nil {
		return err
	}

	r := bytes.NewReader(body)
	links, err := ParseHTML(r)
	if err != nil {
		return err
	}

	visited[u.Path] = struct{}{}
	for _, l := range links {
		childUrl, err := url.Parse(l)
		if err != nil {
			fmt.Println(err)
			continue
		}
		_, isVisited := visited[childUrl.Path]
		if !isVisited && (childUrl.Host == u.Host || childUrl.Host == "") && childUrl.Path != "" {
			fmt.Printf("Inspect\nhost: %s\npath: %s\n------\n", childUrl.Host, childUrl.Path)
			visited[childUrl.Path] = struct{}{}
			//TODO recurse here
			//todo ignore empty path
			if depth < maxDepth || maxDepth == 0 {
				//InspectURL(childUrl, visited, depth + 1, maxDepth)
			}
		} else {
			fmt.Printf("Link %s not inspectable\n------\n", l)
		}
	}

	return nil
}

func ComposeXML(host string, visited map[string]struct{}) ([]byte, error) {
	type XMLUrl struct {
		XMLName xml.Name `xml:"url"`
		Loc     string   `xml:"loc"`
	}
	type UrlSet struct {
		XMLName xml.Name `xml:"urlset"`
		Xmlns   string   `xml:"xmlns,attr"`
		Urls    []XMLUrl `xml:"url"`
	}

	var xmlurls []XMLUrl
	for u, _ := range visited {
		//todo maybe save full url to visited
		xmlurls = append(xmlurls, XMLUrl{Loc : "https://" + host + u})
	}
	us := UrlSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		Urls: xmlurls,
	}

	myString, err := xml.MarshalIndent(us, "", "    ")
	if err != nil {
		return nil, err
	}
	myString = []byte(xml.Header + string(myString))

	return myString, nil
}
