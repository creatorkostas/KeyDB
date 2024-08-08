package db_utils

import (
	database "github.com/creatorkostas/KeyDB/database/database_core/core"
	"github.com/creatorkostas/KeyDB/internal/tools"
)

// database "github.com/creatorkostas/KeyDB/database/database_core"

func SaveDB(filename string) error {
	var err error = nil

	tools.SaveToFile(filename, &database.DB)

	return err
}

func LoadDB(filename string) error {
	var err error = nil

	tools.LoadFromFile(filename, &database.DB)

	return err
}
