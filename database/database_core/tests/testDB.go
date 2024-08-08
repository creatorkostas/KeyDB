package database_test

import (
	"math/rand/v2"
	"sync"
	"time"

	database "github.com/creatorkostas/KeyDB/database/database_core/core"
)

var wg sync.WaitGroup

var tables []string = []string{"table0", "table1", "table2", "table3", "table4", "table5", "table6", "table7", "table8", "table9"}
var keys []string = []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l", "z", "x", "c", "v", "b", "n", "m", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "qq", "ww", "ee", "rr", "tt", "yy", "uu", "ii", "oo", "pp", "aa", "ss", "dd", "ff", "gg", "hh", "jj", "kk", "ll", "zz", "xx", "cc", "vv", "bb", "nn", "mm", "11", "22", "33", "44", "55", "66", "77", "88", "99", "00"}
var tables_num int = 10
var keys_num int = len(keys)
var valus_num int = 5

var values [5]map[string]string

func Run_write_test(concurent_writes int) string {

	for _, v := range tables {
		database.MakeTable(v)
	}

	values[0] = make(map[string]string, 2)
	values[1] = make(map[string]string, 2)
	values[2] = make(map[string]string, 2)
	values[3] = make(map[string]string, 2)
	values[4] = make(map[string]string, 2)

	values[0]["value"] = "22"
	values[0]["type"] = "int"

	values[1]["value"] = "23.34"
	values[1]["type"] = "float64"

	values[2]["value"] = "23.34"
	values[2]["type"] = "float32"

	values[3]["value"] = "true"
	values[3]["type"] = "bool"

	values[4]["value"] = "some string"
	values[4]["type"] = "string"

	now := time.Now()
	for i := 0; i < concurent_writes; i++ {
		wg.Add(1)
		go func() {
			var val_int = rand.IntN(valus_num)
			database.Add_value(tables[rand.IntN(tables_num)], keys[rand.IntN(keys_num)], values[val_int]["type"], values[val_int]["value"])
			wg.Done()
		}()
	}

	wg.Wait()

	return time.Since(now).String()
}

func Run_read_test(concurent_writes int) string {

	// Run_read_test(100)

	now := time.Now()
	for i := 0; i < concurent_writes; i++ {
		wg.Add(1)
		go func() {
			// fmt.Println(
			database.Get_value(tables[rand.IntN(tables_num)], keys[rand.IntN(keys_num)])
			// )
			wg.Done()
		}()
	}

	wg.Wait()

	return time.Since(now).String()
}
