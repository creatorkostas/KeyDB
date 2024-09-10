package tools_test

import (
	"testing"

	"github.com/creatorkostas/KeyDB/internal/tools"
)

func TestSaveFromFile(t *testing.T) {
	var err = tools.SaveToFile("test.gob", "test")
	if err != nil {
		t.Fatalf("SaveToFile failed")
	}
}

func TestLoadFromFile(t *testing.T) {
	var str string
	var err = tools.LoadFromFile("test.gob", &str)
	if err != nil {
		t.Fatalf("LoadFromFile failed")
	}
	if str != "test" {
		t.Fatalf("LoadFromFile faild with corrupted data")
	}
}

func TestDeleteFile(t *testing.T) {
	var err = tools.DeleteFile("test.gob")
	if err != nil {
		t.Fatalf("DeleteFile failed")
	}
}
