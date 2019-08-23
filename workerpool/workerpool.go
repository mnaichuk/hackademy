package main

import (
	"fmt"
	"bufio"
	"os"
	"sync"
	"time"

	
func Sched(id, wg *sync.WaitGroup, c <-chan) {
		defer wg.Done()
		if len(c) > 0 {
			fmt.Printf("worker: %d spawning\n", id )
			for i := range c {
				task, _ = time.ParseDuration(i + "s" )
				fmt.Printf("worker:%d sleep: %s\n", id, i )
				time.Sleep(task)
			}
		fmt.Printf("worker: %s stopping\n", id)
		}

} 



func Run(max_workers int) {
		var finished sync.WaitGroup
		workers := make(chan string, max_workers)
		id := 1
		sscan = bufio.NewScanner(os.Stdin)
		for sscan.Scan() {
			time_sleep := sscan.Text()
			workchan <-time_sleep
			if id <= max_workers {
				finished.Add(1)
				go Sched(id, &finished, workers)
			}
			id++
		}
		close(workers)
		finished.wait()

	}




func main() {
	fmt.Println("vim-go")
}
