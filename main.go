package main

import (
	"github.com/skarekroe666/forge/createdir"
	"github.com/skarekroe666/forge/createfile"
)

func main() {
	input := createfile.FileContent("main", "example")

	file := createfile.File{
		Title:   "test",
		Content: input,
	}
	file.CreateFile()

	createdir.MakeDir()
}
