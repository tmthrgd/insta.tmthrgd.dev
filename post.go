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

var postTmpl = newTemplate("post.tmpl")

var (
	errNoDisplayURL   = errors.New("unable to find display_url in window._sharedData")
	errPrivateAccount = errors.New("this post belongs to a private Instagram account")
)

var postClient = &http.Client{
	Transport: http.DefaultTransport,
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		if strings.Contains(via[0].URL.Path, "/p/") &&
			!strings.Contains(req.URL.Path, "/p/") {
			return errPrivateAccount
		}

		return nil
	},
}

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

		resp, err := postClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return httpError{resp}
		}

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

		data, err := processJSONData(node)
		if err != nil {
			return err
		}

		if strings.HasSuffix(r.URL.Path, "/json") {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			return json.NewEncoder(w).Encode(&struct {
				PostID  string   `json:"post_id"`
				Caption string   `json:"caption,omitempty"`
				Images  []string `json:"images"`
			}{
				chi.URLParam(r, "postID"),
				data.Caption,
				data.DisplayURLs,
			})
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		return templateExecute(w, postTmpl, &struct {
			PostID string
			*jsonData
		}{
			chi.URLParam(r, "postID"),
			data,
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

func processJSONData(n *html.Node) (*jsonData, error) {
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

	var caption string
	if edges := shortcodeMedia.EdgeMediaToCaption.Edges; len(edges) > 0 {
		caption = edges[0].Node.Text
	}

	switch shortcodeMedia.Typename {
	case "GraphImage":
		if shortcodeMedia.DisplayURL != "" {
			return &jsonData{
				Caption:               caption,
				DisplayURLs:           []string{shortcodeMedia.DisplayURL},
				AccessibilityCaptions: []string{shortcodeMedia.AccessibilityCaption},
			}, nil
		}
	case "GraphSidecar":
		edges := shortcodeMedia.EdgeSidecarToChildren.Edges
		displayURLs := make([]string, 0, len(edges))
		captions := make([]string, 0, len(edges))
		for _, edge := range edges {
			if edge.Node.DisplayURL != "" {
				displayURLs = append(displayURLs, edge.Node.DisplayURL)
				captions = append(captions, edge.Node.AccessibilityCaption)
			}
		}

		if len(displayURLs) > 0 {
			return &jsonData{
				Caption:               caption,
				DisplayURLs:           displayURLs,
				AccessibilityCaptions: captions,
			}, nil
		}
	}

	return nil, errNoDisplayURL
}

type jsonData struct {
	Caption               string
	DisplayURLs           []string
	AccessibilityCaptions []string
}

type instagramGraphSchema struct {
	EntryData struct {
		PostPage []struct {
			GraphQL struct {
				ShortcodeMedia struct {
					Typename              string `json:"__typename"`
					DisplayURL            string `json:"display_url"`
					AccessibilityCaption  string `json:"accessibility_caption"`
					EdgeSidecarToChildren struct {
						Edges []struct {
							Node struct {
								DisplayURL           string `json:"display_url"`
								AccessibilityCaption string `json:"accessibility_caption"`
							}
						}
					} `json:"edge_sidecar_to_children"`
					EdgeMediaToCaption struct {
						Edges []struct {
							Node struct {
								Text string
							}
						}
					} `json:"edge_media_to_caption"`
				} `json:"shortcode_media"`
			} `json:"graphql"`
		}
	} `json:"entry_data"`
}
