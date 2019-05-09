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
<link rel=stylesheet href=https://cdnjs.cloudflare.com/ajax/libs/normalize/8.0.1/normalize.min.css>
<link rel=stylesheet href=/assets/style.css>
<form action="" class=download-form>
<div>
<label for=url>Instagram URL: </label>
<input type=url id=url required>
</div>
<input type=submit value=Download>
</form>
<script defer src=/assets/index.js></script>`

func indexHandler() http.HandlerFunc {
	return handlers.ServeString("index.html", time.Now(), indexPage).ServeHTTP
}
