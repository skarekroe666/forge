package createdir

import (
	"log"
	"os"
)

func MakeDir() {
	input := userInput("testdir")

	err := os.Mkdir(input, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}

func userInput(args string) string {
	return args
}
