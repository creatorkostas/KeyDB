package api

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type responce struct {
	Message string `json:"response"`
}

type DB_value struct {
	Value_type string
	Int        int
	String     string
	Float32    float32
	Float64    float64
	Bool       bool
}

var DB map[string]DB_value = make(map[string]DB_value, 100)

func GetValue(c *gin.Context) {

	var key, _ = c.GetQuery("key")

	var stored_data DB_value = DB[key]
	var value_type = stored_data.Value_type

	switch value_type {
	case "int":
		c.IndentedJSON(http.StatusOK, stored_data.Int)
	case "string":
		c.IndentedJSON(http.StatusOK, stored_data.String)
	case "float32":
		c.IndentedJSON(http.StatusOK, stored_data.Float32)
	case "float64":
		c.IndentedJSON(http.StatusOK, stored_data.Float64)
	case "bool":
		c.IndentedJSON(http.StatusOK, stored_data.Bool)
	}
}

// func SetValues(key string, value_type string, value interface{}) {
func SetValues(c *gin.Context) {

	var key, _ = c.GetQuery("key")

	var value_type, _ = c.GetQuery("type")

	var value, _ = c.GetQuery("value")

	DB_val := DB_value{Value_type: value_type}

	switch value_type {
	case "int":
		DB_val.Int, _ = strconv.Atoi(value)
	case "string":
		DB_val.String = value
	case "float32":
		var float, _ = strconv.ParseFloat(value, 32)
		DB_val.Float32 = float32(float)
	case "float64":
		DB_val.Float64, _ = strconv.ParseFloat(value, 64)
	case "bool":
		DB_val.Bool, _ = strconv.ParseBool(value)
	}

	DB[key] = DB_val
	c.IndentedJSON(http.StatusOK, responce{"ok"})
}

func GetSample(c *gin.Context) {
	// var sample_size, _ = c.GetQuery("size")

	// c.IndentedJSON(http.StatusOK, DB[0])
}

func Save(c *gin.Context) {
	// Create a file for IO
	log.Println("test")
	encodeFile, err := os.Create("DB.gob")
	if err != nil {
		log.Println("err")
		panic(err)
	}

	encoder := gob.NewEncoder(encodeFile)

	if err := encoder.Encode(&DB); err != nil {
		log.Println(err)
		panic(err)
	}
	encodeFile.Close()
	log.Println("test2")

	c.IndentedJSON(http.StatusOK, responce{"ok"})
}

func Load(c *gin.Context) {
	decodeFile, err := os.Open("DB.gob")
	if err != nil {
		panic(err)
	}
	defer decodeFile.Close()

	decoder := gob.NewDecoder(decodeFile)

	// Place to decode into
	// accounts2 := make(map[string]Account)

	// Decode -- We need to pass a pointer otherwise accounts2 isn't modified
	decoder.Decode(&DB)

	c.IndentedJSON(http.StatusOK, responce{"ok"})
}
