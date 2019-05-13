// +build generate

package main

import (
	"log"

	"github.com/shurcooL/vfsgen"
	"tmthrgd.dev/go/insta.tmthrgd.dev/internal/templates"
)

func main() {
	if err := vfsgen.Generate(templates.FileSystem, vfsgen.Options{
		Filename:        "internal/templates/vfsdata.go",
		PackageName:     "templates",
		BuildTags:       "!dev",
		VariableName:    "FileSystem",
		VariableComment: "FileSystem contains project templates.",
	}); err != nil {
		log.Fatalln(err)
	}
}
