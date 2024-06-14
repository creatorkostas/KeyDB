package handlers

import (
	"strconv"
	"sync"
)

var m = sync.RWMutex{}

type DB_value struct {
	Value_type *int8
	Int        *int
	String     *string
	Float32    *float32
	Float64    *float64
	Bool       *bool
}

// type Test_DB_Value struct {
// 	uint64
// }

// 0000 0000 - version
// 0000 - type
// 0000
// 0000 0000 0000 0000
// 0000 0000 0000 0000
// 0000 0000 0000 0000

// type Value struct {
// 	Key   string
// 	Value DB_value
// }

var DB map[string]map[string]DB_value = make(map[string]map[string]DB_value, 100)

func Add_value(user string, key string, value_type string, value string) bool {
	var ret bool = false
	// DB_val := DB_value{Value_type: value_type}
	DB_val := DB_value{nil, nil, nil, nil, nil, nil}

	switch value_type {
	case "int":
		*DB_val.Value_type = 0
		*DB_val.Int, _ = strconv.Atoi(value)
		ret = true
	case "string":
		*DB_val.String = value
		*DB_val.Value_type = 1
		ret = true
	case "float32":
		var float, _ = strconv.ParseFloat(value, 32)
		*DB_val.Float32 = float32(float)
		*DB_val.Value_type = 2
		ret = true
	case "float64":
		*DB_val.Value_type = 3
		*DB_val.Float64, _ = strconv.ParseFloat(value, 64)
		ret = true
	case "bool":
		DB_val.Value_type = new(int8)
		*DB_val.Value_type = 4
		DB_val.Bool = new(bool)
		*DB_val.Bool, _ = strconv.ParseBool(value)
		ret = true
	default:
		ret = false
	}

	if ret {
		if len(DB[user]) == 0 {
			DB[user] = make(map[string]DB_value, 100)
		}
		m.Lock()
		DB[user][key] = DB_val
		m.Unlock()
	}
	return ret
}

func Get_value(user string, key string) any {

	m.RLock()
	var stored_data = DB[user][key]
	m.RUnlock()
	var value_type = stored_data.Value_type
	if value_type == nil {
		return nil
	}
	var value any

	switch *value_type {
	case 0:
		value = *stored_data.Int
	case 1:
		value = *stored_data.String
	case 2:
		value = *stored_data.Float32
	case 3:
		value = *stored_data.Float64
	case 4:
		value = *stored_data.Bool
	default:
		value = nil
	}

	return value
}

func Get_Users_Data(user string) map[string]DB_value {
	return DB[user]
}
