package main

import (
	"net/http"
	urlpkg "net/url"
	"os"
	"strings"
)

func pRedirectHandler() http.HandlerFunc {
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		if err := r.ParseForm(); err != nil {
			return err
		}

		url := r.Form.Get("url")
		if url == "" {
			return os.ErrNotExist
		}

		if path, ok := validInstagramURL(url); ok {
			http.Redirect(w, r, path, http.StatusSeeOther)
			return nil
		}

		http.Redirect(w, r, "/", http.StatusFound)
		return nil
	})
}

func validInstagramURL(url string) (string, bool) {
	u, err := urlpkg.Parse(url)
	if err != nil {
		return "", false
	}

	switch u.Scheme {
	case "http", "https":
	default:
		return "", false
	}

	switch strings.ToLower(u.Host) {
	case "instagram.com", "www.instagram.com", "m.instagram.com":
	default:
		return "", false
	}

	pIdx := strings.Index(u.Path, "/p/")
	if pIdx < 0 {
		return "", false
	}

	postID := strings.TrimSuffix(u.Path[pIdx+len("/p/"):], "/")
	if postID == "" {
		return "", false
	}

	for _, r := range postID {
		if !validBase64(r) {
			return "", false
		}
	}

	path := u.Path[pIdx:]
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	return path, true
}

func validBase64(r rune) bool {
	return r >= 'A' && r <= 'Z' ||
		r >= 'a' && r <= 'z' ||
		r >= '0' && r <= '9' ||
		r == '-' || r == '_'
}
