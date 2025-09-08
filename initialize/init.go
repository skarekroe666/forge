package initialize

import (
	"log"

	"github.com/skarekroe666/forge/createdir"
	"github.com/skarekroe666/forge/makefile"
)

type File struct {
	Title   string
	Content string
}

func InitFileDir() {
	fileTitle := createdir.UserInput("testdir")
	input := makefile.FileContent("main", "example")

	createdir.MakeDir(fileTitle)

	file := File{
		Title:   fileTitle,
		Content: input,
	}

	if err := makefile.CreateFile(file.Title, file.Content); err != nil {
		log.Fatal(err)
	}
}
