package database

import (
	"sync"
)

const (
	INT            = "int"
	STRING         = "string"
	FLOAT32        = "float32"
	FLOAT64        = "float64"
	BOOL           = "bool"
	ENCRYPTED_DATA = "encrypted_data"
)

const (
	INT_INT            = 0
	INT_STRING         = 1
	INT_FLOAT32        = 2
	INT_FLOAT64        = 3
	INT_BOOL           = 4
	INT_ENCRYPTED_DATA = 5
)

var m = sync.RWMutex{}

type DB_value struct {
	Value_type int8
	Data       []byte
}

func makeDefault_DB_value() DB_value {
	return DB_value{-1, nil}
}

var DB map[string]map[string]DB_value = make(map[string]map[string]DB_value, 100)
