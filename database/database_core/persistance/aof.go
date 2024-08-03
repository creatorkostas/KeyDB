package persistance

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"sync"

	internal "github.com/creatorkostas/KeyDB/database/database_core/conf"
)

var wg sync.WaitGroup
var Operations = make(chan string, 256)

func Writer(msg string) {
	Operations <- msg + "\n"
}

func Reader(channel <-chan string) {
	msg := <-channel
	fmt.Println(msg)
}

func write(write_operations [internal.Append_size]string) error {
	var s strings.Builder
	for _, v := range write_operations {
		s.WriteString(v)
	}
	var file, err = os.OpenFile(internal.Append_file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, fs.ModeAppend)
	file.WriteString(s.String())
	file.Close()

	return err
}

func add_to_file(instance int) {
	log.Println(fmt.Sprint("Starting writter ", instance+1))
	var write_operations [internal.Append_size]string
	var i int = 0
	var run bool = true
	for run {
		write_operations[i] = <-Operations
		// log.Println(write_operations[i])
		// log.Println(i)
		if write_operations[i] == "||exit||" {
			run = false
			// break
		}
		i++
		// log.Println(i)

		if i == internal.Append_size {
			write(write_operations)
			i = 0
		}
		// log.Println(i)
	}

	write(write_operations)
	defer wg.Done()
}

func Start_writers(instances int) {
	for i := 0; i < instances; i++ {
		wg.Add(1)
		go add_to_file(i)
	}

	go wg.Wait()
}
