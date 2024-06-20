package database

import (
	"sync"
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
