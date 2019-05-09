package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html/charset"
)

var postTmpl = newTemplate(`<!doctype html>
<html lang=en>
<meta charset=utf-8>
<meta name=viewport content="width=device-width,initial-scale=1">
<title>{{.PostID}} – insta.tmthrgd.dev</title>
<link rel=stylesheet href=https://cdnjs.cloudflare.com/ajax/libs/normalize/8.0.1/normalize.min.css>
<link rel=stylesheet href=/assets/style.css>
<ul class=images>
{{- range .DisplayURLs}}
<li><img src="{{.}}"></li>
{{- end}}
</ul>`)

var errNoDisplayURL = errors.New("unable to find display_url in window._sharedData")

func postHandler() http.HandlerFunc {
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		u := &url.URL{
			Scheme: "https",
			Host:   "www.instagram.com",
			Path:   r.URL.Path,
		}
		// TODO: use http.NewRequestWithContext once go1.13 lands.
		req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
		req = req.WithContext(r.Context())

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
		if err != nil {
			return err
		}

		doc, err := html.Parse(body)
		if err != nil {
			return err
		}

		node, err := findJSONData(doc)
		if err != nil {
			return err
		}

		displayURLs, err := processJSONData(node)
		if err != nil {
			return err
		}

		return templateExecute(w, postTmpl, &struct {
			PostID      string
			DisplayURLs []string
		}{
			chi.URLParam(r, "postID"),
			displayURLs,
		})
	})
}

func findJSONData(doc *html.Node) (*html.Node, error) {
	var (
		node *html.Node
		walk func(*html.Node) bool
	)
	walk = func(n *html.Node) bool {
		if nodeIsJSONData(n) {
			node = n
			return false
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if !walk(c) {
				return false
			}
		}

		return true
	}
	if walk(doc) {
		return nil, errors.New("_sharedData JSON was not found in HTML")
	}

	return node, nil
}

func nodeIsJSONData(n *html.Node) bool {
	return n.Type == html.ElementNode &&
		n.DataAtom == atom.Script &&
		n.FirstChild != nil &&
		n.FirstChild.NextSibling == nil &&
		strings.HasPrefix(n.FirstChild.Data, "window._sharedData")
}

func processJSONData(n *html.Node) ([]string, error) {
	data := n.FirstChild.Data
	data = strings.TrimLeftFunc(data, func(r rune) bool {
		return r != '{' // everything until the opening bracket
	})
	data = strings.TrimRightFunc(data, func(r rune) bool {
		return r != '}' // everything after the closing bracket
	})

	var instData instagramGraphSchema
	if err := json.Unmarshal([]byte(data), &instData); err != nil {
		return nil, err
	}

	if len(instData.EntryData.PostPage) == 0 {
		return nil, errNoDisplayURL
	}

	shortcodeMedia := instData.EntryData.PostPage[0].GraphQL.ShortcodeMedia
	switch shortcodeMedia.Typename {
	case "GraphImage":
		if shortcodeMedia.DisplayURL != "" {
			return []string{shortcodeMedia.DisplayURL}, nil
		}
	case "GraphSidecar":
		edges := shortcodeMedia.EdgeSidecarToChildren.Edges
		displayURLs := make([]string, 0, len(edges))
		for _, edge := range edges {
			if edge.Node.DisplayURL != "" {
				displayURLs = append(displayURLs, edge.Node.DisplayURL)
			}
		}

		if len(displayURLs) > 0 {
			return displayURLs, nil
		}
	}

	return nil, errNoDisplayURL
}

type instagramGraphSchema struct {
	EntryData struct {
		PostPage []struct {
			GraphQL struct {
				ShortcodeMedia struct {
					Typename              string `json:"__typename"`
					DisplayURL            string `json:"display_url"`
					EdgeSidecarToChildren struct {
						Edges []struct {
							Node struct {
								DisplayURL string `json:"display_url"`
							}
						}
					} `json:"edge_sidecar_to_children"`
				} `json:"shortcode_media"`
			} `json:"graphql"`
		}
	} `json:"entry_data"`
}
