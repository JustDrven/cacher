package manager

import (
	"log"
	"os"
)

func CheckFiles() {
	for _, value := range GetFileList() {
		_, err := os.Open(value)
		if err != nil {

			log.Fatalln("File '" + value + "' doesn't exist!")
			break
		}
	}
}

func GetFileList() []string {
	return []string{".env", "go.mod", "start.sh"}
}
