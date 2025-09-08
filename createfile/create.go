package createfile

import (
	"fmt"
	"os"
)

type File struct {
	Title   string
	Content string
}

func (f *File) CreateFile() error {
	fileName := f.Title + ".go"
	return os.WriteFile(fileName, []byte(f.Content), 0644)
}

func FileContent(pkgName, funcName string) string {
	line1 := "package " + pkgName
	line2 := fmt.Sprintf("func %s() {}", funcName)

	lines := []string{line1, line2}
	var content string

	for _, line := range lines {
		content += fmt.Sprintf("%s\n\n", line)
	}

	return content
}
