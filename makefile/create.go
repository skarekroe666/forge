package makefile

import (
	"errors"
	"fmt"
	"os"
)

func CreateFile(title, content string) error {
	fileName := title + ".go"
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return errors.New("failed to write to the file")
	}

	return nil
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
