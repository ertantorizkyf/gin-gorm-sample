package helpers

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadJson(filePath string) []byte {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("[ERR] Failed to read json file")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}
