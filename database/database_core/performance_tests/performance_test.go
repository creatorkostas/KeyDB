package performance_tests

import (
	"math/rand/v2"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	database "github.com/creatorkostas/KeyDB/database/database_core/core"
	"github.com/jedib0t/go-pretty/table"
)

var wg sync.WaitGroup
var concurent_writes int = 10000
var concurent_reads int = 10000

var tables []string = []string{"table1", "table2", "table3", "table4", "table5", "table6", "table7", "table8", "table9", "table10"}
var keys []string = []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l", "z", "x", "c", "v", "b", "n", "m", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "qq", "ww", "ee", "rr", "tt", "yy", "uu", "ii", "oo", "pp", "aa", "ss", "dd", "ff", "gg", "hh", "jj", "kk", "ll", "zz", "xx", "cc", "vv", "bb", "nn", "mm", "11", "22", "33", "44", "55", "66", "77", "88", "99", "00"}
var tables_num int = len(tables)
var keys_num int = len(keys)
var valus_num int = 5

var values [5]map[string]string

func TestWriteSpeed(t *testing.T) {

	c_table := table.NewWriter()
	c_table.SetOutputMirror(os.Stdout)
	c_table.AppendHeader(table.Row{"time", "type"})

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

	values[1]["value"] = "23.345"
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
			database.Add_value(tables[rand.IntN(tables_num)], keys[rand.IntN(keys_num)], values[val_int]["type"], values[val_int]["value"], false, false, "")
			wg.Done()
		}()
	}

	wg.Wait()

	var total_time = time.Since(now).Milliseconds()
	c_table.AppendRows([]table.Row{{total_time, "ms"}})

	if total_time > 2000 {
		c_table.Render()
		t.Fatalf("[WARNING]: Write speed test took too long (" + strconv.Itoa(int(total_time)) + " ms)")
	}

	c_table.Render()

}

func TestReadSpeed(t *testing.T) {

	// Run_read_test(100)
	c_table := table.NewWriter()
	c_table.SetOutputMirror(os.Stdout)
	c_table.AppendHeader(table.Row{"time", "type"})

	now := time.Now()
	for i := 0; i < concurent_reads; i++ {
		wg.Add(1)
		go func() {
			database.Get_value(tables[rand.IntN(tables_num)], keys[rand.IntN(keys_num)], false)
			wg.Done()
		}()
	}

	wg.Wait()

	var total_time = time.Since(now).Milliseconds()
	c_table.AppendRows([]table.Row{{total_time, "ms"}})

	if total_time > 2000 {
		c_table.Render()
		t.Fatalf("[WARNING]: Read speed test took too long (" + strconv.Itoa(int(total_time)) + " ms)")
	}

	c_table.Render()
}
