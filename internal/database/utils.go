package database

import (
	"bytes"
	"encoding/binary"
	"log"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FF
// FF
// FF
// FF
// FF
// FF
// FF
// F

func Add_value(user string, key string, value_type string, c *gin.Context) bool {
	var ret bool = false
	DB_val := makeDefault_DB_value()

	// var temp any

	switch value_type {
	case "int":
		DB_val.Value_type = 0
		var data, _ = c.GetQuery("value")
		var d, _ = strconv.Atoi(data)
		DB_val.Data = make([]byte, 8)
		for i := 0; i < 8; i++ {
			DB_val.Data[i] = byte(int64(d) >> (i * 8) & 0xFF)
		}
	case "string":
		DB_val.Value_type = 1
		DB_val.Data = []byte(c.GetString("value"))
	case "float32":
		DB_val.Value_type = 2

		DB_val.Data = make([]byte, 4)

		var data, _ = c.GetQuery("value")
		float, err := strconv.ParseFloat(data, 32)

		if err != nil {
			log.Println(err)
		}

		n := math.Float32bits(float32(float))

		for i := 0; i < 4; i++ {
			DB_val.Data[i] = byte(n >> (i * 8))
		}

	case "float64":
		DB_val.Value_type = 3
		DB_val.Data = make([]byte, 8)

		var data, _ = c.GetQuery("value")
		float, err := strconv.ParseFloat(data, 64)

		if err != nil {
			log.Println(err)
		}

		n := math.Float64bits(float)

		for i := 0; i < 8; i++ {
			DB_val.Data[i] = byte(n >> (i * 8))
		}
	case "bool":
		DB_val.Value_type = 4
		DB_val.Data = make([]byte, 1)
		var value, _ = c.GetQuery("value")
		if value == "0" || value == "false" {
			DB_val.Data[0] = 0x00
		} else if value == "1" || value == "true" {
			DB_val.Data[0] = 0x01
		}

	default:
		ret = false
	}

	ret = true

	if ret {
		// MakeUserDB(user)
		// log.Println(DB)
		m.Lock()
		DB[user][key] = DB_val
		m.Unlock()
	}
	return ret
}

func Get_value(user string, key string) any {

	if key == "user.get.all.data" {
		m.RLock()
		var stored_data = DB[user]
		m.RUnlock()
		return stored_data
	} else {

		m.RLock()
		var stored_data = DB[user][key]
		m.RUnlock()
		var value_type = stored_data.Value_type
		if value_type == -1 {
			return nil
		}

		switch value_type {
		case 0:
			var data, err = dataConvert[int64](stored_data.Data)
			if err != nil {
				log.Println(err)
				return nil
			}
			return data
		case 1:
			var data = string(stored_data.Data)
			return data
		case 2:
			var data, err = dataConvert[float32](stored_data.Data)
			if err != nil {
				log.Println(err)
				return nil
			}
			return data
		case 3:
			var data, err = dataConvert[float64](stored_data.Data)
			if err != nil {
				log.Println(err)
				return nil
			}
			return data
		case 4:
			var data, err = dataConvert[bool](stored_data.Data)
			if err != nil {
				log.Println(err)
				return nil
			}
			return data
		default:
			return nil
		}
	}
}

func dataConvert[T int | int64 | int32 | float32 | float64 | bool | string](data []byte) (T, error) {
	var temp T
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.LittleEndian, &temp)
	return temp, err
}

func MakeUserDB(user string) {
	if len(DB[user]) == 0 {
		DB[user] = make(map[string]DB_value, 100)
	}
}
