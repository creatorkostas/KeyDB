package database_test

import (
	"os"
	"reflect"
	"testing"

	database "github.com/creatorkostas/KeyDB/database/database_core/core"
	"github.com/jedib0t/go-pretty/table"
)

// ---------------------------| int         | string          | float32       | float64          | bool
var types []string = []string{database.INT, database.STRING, database.FLOAT32, database.FLOAT64, database.BOOL} // "encrypted_data"}
var values []string = []string{"2", "some string", "23.34", "23.345", "true"}                                   //"some string"}
var keys []string = []string{"key0", "key1", "key2", "key3", "key4"}                                            //"key5"}

func TestAddValue(t *testing.T) {
	database.MakeTable("table0")

	c_table := table.NewWriter()
	c_table.SetOutputMirror(os.Stdout)
	c_table.AppendHeader(table.Row{"key", "type", "value"})

	var err error

	for key, value := range keys {
		c_table.AppendRows([]table.Row{
			{value, types[key], values[key]},
		})
		err = database.Add_value("table0", value, types[key], values[key], false, false, "")
		if err != nil {
			t.Fatalf("[ERROR]: " + err.Error())
		}
	}

	c_table.Render()

}

func TestGetValue(t *testing.T) {

	var data any

	c_table := table.NewWriter()
	c_table.SetOutputMirror(os.Stdout)
	c_table.AppendHeader(table.Row{"key", "type", "value"})

	for _, value := range keys {
		data = database.Get_value("table0", value, false)

		if data == nil {
			c_table.Render()
			t.Fatalf("[ERROR]: Data is nil")
		} else {
			switch data := data.(type) {
			case int:
				if data != 2 {
					c_table.AppendRows([]table.Row{{value, reflect.TypeOf(data), data}})
					c_table.Render()
					t.Fatalf("[ERROR]: Data is not 2")
				}
				c_table.AppendRows([]table.Row{{value, reflect.TypeOf(data), data}})
			case int64:
				if data != 2 {
					c_table.AppendRows([]table.Row{{value, reflect.TypeOf(data), data}})
					c_table.Render()
					t.Fatalf("[ERROR]: Data is not 2")
				}
				c_table.AppendRows([]table.Row{{value, reflect.TypeOf(data), data}})
			case int32:
				if data != 2 {
					c_table.AppendRows([]table.Row{{value, reflect.TypeOf(data), data}})
					c_table.Render()
					t.Fatalf("[ERROR]: Data is not 2")
				}
				c_table.AppendRows([]table.Row{{value, reflect.TypeOf(data), data}})
			case string:
				if data != "some string" {
					c_table.AppendRows([]table.Row{{value, reflect.TypeOf(data), data}})
					c_table.Render()
					t.Fatalf("[ERROR]: Data is not 'some string'")
				}
				c_table.AppendRows([]table.Row{{value, reflect.TypeOf(data), data}})
			case float32:
				if data != 23.34 {
					c_table.AppendRows([]table.Row{{value, reflect.TypeOf(data), data}})
					c_table.Render()
					t.Fatalf("[ERROR]: Data is not 23.34")
				}
				c_table.AppendRows([]table.Row{{value, reflect.TypeOf(data), data}})
			case float64:
				if data != 23.345 {
					c_table.AppendRows([]table.Row{{value, reflect.TypeOf(data), data}})
					c_table.Render()
					t.Fatalf("[ERROR]: Data is not 23.345")
				}
				c_table.AppendRows([]table.Row{{value, reflect.TypeOf(data), data}})
			case bool:
				if data != true {
					c_table.AppendRows([]table.Row{{value, reflect.TypeOf(data), data}})
					c_table.Render()
					t.Fatalf("[ERROR]: Data is not true")
				}
				c_table.AppendRows([]table.Row{{value, reflect.TypeOf(data), data}})
			default:
				c_table.AppendRows([]table.Row{{value, reflect.TypeOf(data), data}})
				c_table.Render()
				t.Fatalf("[ERROR]: Data type is not supported")
			}

		}
	}

}
