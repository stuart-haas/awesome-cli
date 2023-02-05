package utils

import (
	"log"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return data, err
}
