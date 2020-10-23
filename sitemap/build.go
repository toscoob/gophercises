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
	"strings"
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

func hostMatch(baseUrl url.URL, checkUrl url.URL) bool {
	return checkUrl.Host == baseUrl.Host || checkUrl.Host == "" || checkUrl.Host == "www." + baseUrl.Host
}

func emptyPath(u url.URL) bool {
	return strings.TrimSpace(u.Path) == "" || u.Path == "/"
}

func InspectURL(u url.URL, visited map[string]struct{}, depth uint, maxDepth uint) error {
	fmt.Printf("###InspectURL depth %d ###\n", depth)
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
		isVisited = isVisited || emptyPath(*childUrl)
		sameHost := hostMatch(u, *childUrl)
		if !isVisited && sameHost {
			fmt.Printf("Inspect and add \nhost: %s path: %s\n------\n", childUrl.Host, childUrl.Path)
			//visited[childUrl.Path] = struct{}{}
			//TODO recurse here
			//todo ignore empty path
			if depth < maxDepth || maxDepth == 0 {
				childUrl.Scheme = u.Scheme
				childUrl.Host = u.Host //in case it's relative link
				err = InspectURL(*childUrl, visited, depth + 1, maxDepth)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
		} else {
			var reason string
			switch {
			case isVisited:
				reason = "visited"
			case emptyPath(*childUrl):
				reason = "empty path"
			case !sameHost:
				reason = "different host"
			default:
				reason = "unknown"
			}
			fmt.Printf("Link %s not inspectable because %s\n------\n", l, reason)
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
	for u := range visited {
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
