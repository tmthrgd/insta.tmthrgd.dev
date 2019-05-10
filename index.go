package main

import (
	"net/http"
	"time"

	handlers "github.com/tmthrgd/httphandlers"
)

const indexPage = `<!doctype html>
<html lang=en>
<meta charset=utf-8>
<meta name=viewport content="width=device-width,initial-scale=1">
<title>insta.tmthrgd.dev</title>
<link rel=stylesheet href=https://cdnjs.cloudflare.com/ajax/libs/normalize/8.0.1/normalize.min.css integrity="sha256-l85OmPOjvil/SOvVt3HnSSjzF1TUMyT9eV0c2BzEGzU=" crossorigin=anonymous>
<link rel=stylesheet href=https://cdnjs.cloudflare.com/ajax/libs/skeleton/2.0.4/skeleton.min.css integrity="sha256-2YQRJMXD7pIAPHiXr0s+vlRWA7GYJEK0ARns7k2sbHY=" crossorigin=anonymous>
<link rel=stylesheet href="https://fonts.googleapis.com/css?family=Raleway">
<link rel=stylesheet href=/assets/index.css>
<main class=container>
<form action="" class=download-form>
<div class=row>
<label for=url>Instagram URL: </label>
<input type=url class=u-full-width id=url required placeholder=https://www.instagram.com/p/... pattern="https?://(www\.|m\.)?instagram\.com/(.+/)?p/[A-Za-z0-9\-_]+/?(\?.*)?(#.*)?">
</div>
<input type=submit value=Download>
</form>
</main>
<script defer src=/assets/index.js></script>`

func indexHandler() http.HandlerFunc {
	return handlers.ServeString("index.html", time.Now(), indexPage).ServeHTTP
}
