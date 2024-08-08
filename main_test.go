package main

import (
	"testing"

	database "github.com/creatorkostas/KeyDB/database/database_core/core"
	// "github.com/creatorkostas/KeyDB/internal/database/database_test"
)

// var DB_filename string = "db.gob"

// func TestSpeed(_ *testing.T) {
// 	var num_of_reads int = 1000
// 	var num_of_writes int = 100

// 	fmt.Println("Concurent writes (" + strconv.Itoa(num_of_writes) + ") : " + database_test.Run_write_test(num_of_writes))
// 	fmt.Println("Concurent reads (" + strconv.Itoa(num_of_reads) + ") : " + database_test.Run_read_test(num_of_reads))

// }

func TestWriteRead_INT(t *testing.T) {
	// var get_value any
	database.MakeTable("test")

	database.Add_value("test", "test", "int", "2")
	// get_value = database.Get_value("test", "test")

	// fmt.Println(get_value)
	// fmt.Println(database.DB)
	// fmt.Println(database.Get_value("test", "table.get.all.data"))
	// if get_value != 2 {
	t.Failed()
	// }

}
