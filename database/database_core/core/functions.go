package database

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"unsafe"

	"github.com/creatorkostas/KeyDB/database/database_core/persistance"
	"github.com/creatorkostas/KeyDB/database/database_core/security"
	// "github.com/creatorkostas/KeyDB/database/database_core/persistance"
)

func Add_value(table string, key string, value_type string, data string, encrypted_data bool, encrypt_on_save bool, public_key string) error {
	// var ret bool = false
	// var temp_data = data
	//TODO  see encryption keys
	if encrypted_data && public_key != "" {
		data = security.Decrypt_data(public_key, []byte(data))
	}

	var set_err error
	// if !internal.Append_only_in_file {
	DB_val := makeDefault_DB_value()

	if encrypt_on_save && public_key != "" {
		DB_val.Value_type = INT_ENCRYPTED_DATA
		DB_val.Data = security.Encrypt_data(public_key, []byte(data))
	} else {

		switch value_type {
		case INT:
			DB_val.Value_type = INT_INT
			var d, err = strconv.Atoi(data)
			set_err = err

			if err == nil {
				DB_val.Data = make([]byte, 8)
				for i := 0; i < 8; i++ {
					DB_val.Data[i] = byte(int64(d) >> (i * 8) & 0xFF)
				}
			}

		case STRING:
			DB_val.Value_type = INT_STRING
			DB_val.Data = []byte(data)
		case FLOAT32:
			DB_val.Value_type = INT_FLOAT32
			DB_val.Data = make([]byte, 8)
			float, err := strconv.ParseFloat(data, 32)
			set_err = err

			float32bit := float32(float)
			if err == nil {
				n := *(*uint32)(unsafe.Pointer(&float32bit))

				for i := 0; i < 8; i++ {
					DB_val.Data[i] = byte(n >> (i * 8))
				}
			}

		case FLOAT64:
			DB_val.Value_type = INT_FLOAT64
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
			DB_val.Value_type = INT_BOOL
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

func Get_value(table string, key string, get_raw bool) any {

	if key == "table.get.all.data" {
		m.RLock()
		var stored_data = DB[table]
		m.RUnlock()
		return stored_data
	} else if get_raw {
		m.RLock()
		var stored_data = DB[table][key]
		m.RUnlock()
		return stored_data.Data
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
			return data
		case INT_STRING:
			var data = string(stored_data.Data)
			return data
		case INT_FLOAT32:
			// Is still saved as float64. It needs to be changes in the AddValue function
			var data, _ = dataConvert[float32](stored_data.Data)
			return data
		case INT_FLOAT64:
			var data, _ = dataConvert[float64](stored_data.Data)
			return data
		case INT_BOOL:
			return *(*bool)(unsafe.Pointer(&stored_data.Data))
		case INT_ENCRYPTED_DATA:
			return stored_data.Data
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

func DeleteTable(table string) {
	delete(DB, table)
}
