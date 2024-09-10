package db_utils_test

import (
	"testing"

	db_utils "github.com/creatorkostas/KeyDB/database/database_core/utils"
)

func TestSaveDB(t *testing.T) {
	t.SkipNow()
	var err = db_utils.SaveDB("test.gob")
	if err != nil {
		t.Fatalf("SaveDB failed")
	}
}

func TestLoadDB(t *testing.T) {
	var err = db_utils.LoadDB("test.gob")
	if err != nil {
		t.Fatalf("LoadDB failed")
	}
}
