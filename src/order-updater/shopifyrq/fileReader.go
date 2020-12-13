package shopifyrq

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func FileReader(fileDirectory string) []byte {

	// Open our jsonFile
	jsonFile, err := os.Open(fileDirectory)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	log.Println("Successfully Opened the file.")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}
