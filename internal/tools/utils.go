package tools

import (
	"encoding/gob"
	"log"
	"os"

	"github.com/creatorkostas/KeyDB/api"
)

func Save() {
	// Create a file for IO
	log.Println("test")
	encodeFile, err := os.Create("db.gob")
	if err != nil {
		log.Println("err")
		panic(err)
	}

	encoder := gob.NewEncoder(encodeFile)

	if err := encoder.Encode(&api.DB); err != nil {
		log.Println(err)
		panic(err)
	}
	encodeFile.Close()
	log.Println("test2")

}

func Load() {
	decodeFile, err := os.Open("db.gob")
	if err != nil {
		panic(err)
	}
	defer decodeFile.Close()

	decoder := gob.NewDecoder(decodeFile)

	// Place to decode into
	// accounts2 := make(map[string]Account)

	// Decode -- We need to pass a pointer otherwise accounts2 isn't modified
	decoder.Decode(&api.DB)

}
