package createdir

import (
	"log"
	"os"
)

func MakeDir(input string) {
	err := os.Mkdir(input, 0750)
	if err != nil && os.IsNotExist(err) {
		log.Fatal(err)
	}
}

func UserInput(args string) string {
	return args
}

// func createMultiDir() {

// }
