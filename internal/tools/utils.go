package tools

import (
	"encoding/gob"
	"log"
	"os"
)

func SaveDB(filename string, data any) {
	// Create a file for IO
	encodeFile, err := os.Create(filename)
	if err != nil {
		log.Println("err")
		panic(err)
	}

	encoder := gob.NewEncoder(encodeFile)

	if err := encoder.Encode(data); err != nil {
		log.Println(err)
		panic(err)
	}
	encodeFile.Close()

}

func LoadDB(filename string, data any) {
	decodeFile, err := os.Open(filename)
	// decodeFile, err := os.Open("db.gob")
	if err != nil {
		log.Println("Possible the file does not exist!!")
		// panic(err)
	}
	defer decodeFile.Close()

	decoder := gob.NewDecoder(decodeFile)

	// Place to decode into
	// accounts2 := make(map[string]Account)

	// Decode -- We need to pass a pointer otherwise accounts2 isn't modified
	decoder.Decode(data)

}
