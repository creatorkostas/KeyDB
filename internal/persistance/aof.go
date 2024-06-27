package aof

import (
	"io/fs"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var Operations = make(chan string, 256)

func write(write_operations [64]string) error {
	var s strings.Builder
	var file, err = os.OpenFile("", os.O_APPEND, fs.ModeAppend)
	for _, v := range write_operations {
		s.WriteString(v)
	}
	file.WriteString(s.String())
	file.Close()

	return err
}

func add_to_file() {
	var write_operations [64]string
	var i int = 0
	var run bool = true
	for run {
		write_operations[i] = <-Operations
		if write_operations[i] == "||exit||" {
			run = false
			break
		}
		i++

		if len(write_operations) == 64 {
			write(write_operations)
			i = 0
		}
	}

	write(write_operations)
	defer wg.Done()
}

func Start_writers(instances int) {
	// for i := 0; i < instances; i++ {
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go add_to_file()
	}

	wg.Wait()
}
