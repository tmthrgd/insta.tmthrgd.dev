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
<link rel=stylesheet href=https://cdnjs.cloudflare.com/ajax/libs/normalize/8.0.1/normalize.min.css integrity="sha256-l85OmPOjvil/SOvVt3HnSSjzF1TUMyT9eV0c2BzEGzU=" crossorigin=anonymous>
<link rel=stylesheet href=https://cdnjs.cloudflare.com/ajax/libs/skeleton/2.0.4/skeleton.min.css integrity="sha256-2YQRJMXD7pIAPHiXr0s+vlRWA7GYJEK0ARns7k2sbHY=" crossorigin=anonymous>
<link rel=stylesheet href="https://fonts.googleapis.com/css?family=Raleway">
<link rel=stylesheet href=/assets/post.css>
<main class=container>
{{- if eq (len .DisplayURLs) 1}}
<img class=post-image src="{{index .DisplayURLs 0}}">
{{- else}}
<ul class=post-images>
{{- range .DisplayURLs}}
<li><img class=post-image src="{{.}}"></li>
{{- end}}
</ul>
{{- end}}
</main>`)

var errNoDisplayURL = errors.New("unable to find display_url in window._sharedData")

func postHandler() http.HandlerFunc {
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		path := r.URL.Path
		if strings.HasSuffix(path, "/json") {
			path = path[:len(path)-len("json")]
		}

		u := &url.URL{
			Scheme: "https",
			Host:   "www.instagram.com",
			Path:   path,
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

		if strings.HasSuffix(r.URL.Path, "/json") {
			return json.NewEncoder(w).Encode(&struct {
				PostID string   `json:"post_id"`
				Images []string `json:"images"`
			}{
				chi.URLParam(r, "postID"),
				displayURLs,
			})
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
