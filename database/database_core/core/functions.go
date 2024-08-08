package database

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"unsafe"

	"github.com/creatorkostas/KeyDB/database/database_core/persistance"
	// "github.com/creatorkostas/KeyDB/database/database_core/persistance"
)

func Add_value(table string, key string, value_type string, data string) error {
	// var ret bool = false
	var set_err error
	// if !internal.Append_only_in_file {
	DB_val := makeDefault_DB_value()

	switch value_type {
	case INT:
		DB_val.Value_type = 0
		var d, err = strconv.Atoi(data)
		set_err = err

		if err == nil {
			DB_val.Data = make([]byte, 8)
			for i := 0; i < 8; i++ {
				DB_val.Data[i] = byte(int64(d) >> (i * 8) & 0xFF)
			}
		}
		// fmt.Println("aaaaa")
		// fmt.Println(DB_val.Data)
		// fmt.Println(err)

	case STRING:
		DB_val.Value_type = 1
		DB_val.Data = []byte(data)
	case FLOAT32:
		// TODO make 32 bit
		DB_val.Value_type = 3
		DB_val.Data = make([]byte, 8)
		float, err := strconv.ParseFloat(data, 64)
		set_err = err

		if err == nil {
			n := *(*uint64)(unsafe.Pointer(&float))

			for i := 0; i < 8; i++ {
				DB_val.Data[i] = byte(n >> (i * 8))
			}
		}

	case FLOAT64:
		DB_val.Value_type = 3
		DB_val.Data = make([]byte, 8)

		float, err := strconv.ParseFloat(data, 64)
		set_err = err

		if err == nil {
			n := *(*uint64)(unsafe.Pointer(&float))

			for i := 0; i < 8; i++ {
				DB_val.Data[i] = byte(n >> (i * 8))
			}
		}

	case BOOL:
		DB_val.Value_type = 4
		DB_val.Data = make([]byte, 1)

		if data == "0" || data == "false" {
			DB_val.Data[0] = 0x00
		} else if data == "1" || data == "true" {
			DB_val.Data[0] = 0x01
		}

	default:
		// ret = false
		set_err = errors.New("something went wrong")
	}

	if set_err == nil {
		m.Lock()
		DB[table][key] = DB_val
		m.Unlock()
	}
	// }

	go func() {
		persistance.Writer(fmt.Sprintf("%s|%s|%s|%s", table, key, value_type, data))
	}()

	return set_err
}

func Get_value(table string, key string) any {

	if key == "table.get.all.data" {
		m.RLock()
		var stored_data = DB[table]
		m.RUnlock()
		return stored_data
	} else {

		m.RLock()
		var stored_data = DB[table][key]
		m.RUnlock()
		var value_type = stored_data.Value_type
		if value_type == -1 {
			return nil
		}

		switch value_type {
		case INT_INT:
			var data, _ = dataConvert[int64](stored_data.Data)
			// fmt.Println(*(*int64)(unsafe.Pointer(&stored_data.Data)))
			// if err != nil {
			// 	log.Println(err)
			// 	return nil
			// }
			return data
			// return *(*int64)(unsafe.Pointer(&stored_data.Data))
		case INT_STRING:
			var data = string(stored_data.Data)
			return data
		case INT_FLOAT32:
			var data, _ = dataConvert[float32](stored_data.Data)
			// if err != nil {
			// 	log.Println(err)
			// 	return nil
			// }
			return data
			// return *(*float32)(unsafe.Pointer(&stored_data.Data))
		case INT_FLOAT64:
			var data, _ = dataConvert[float64](stored_data.Data)
			// if err != nil {
			// 	log.Println(err)
			// 	return nil
			// }
			return data
			// return *(*float64)(unsafe.Pointer(&stored_data.Data))
		case INT_BOOL:
			// var data, err = dataConvert[bool](stored_data.Data)
			// if err != nil {
			// 	log.Println(err)
			// 	return nil
			// }
			// return data

			return *(*bool)(unsafe.Pointer(&stored_data.Data))
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

func MakeTable(table string) {
	if len(DB[table]) == 0 {
		DB[table] = make(map[string]DB_value, 100)
	}
}
