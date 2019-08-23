package pool

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Worker struct {
	id     int
	status bool
}

func Job(pool chan float64, worker Worker, wg *sync.WaitGroup) {
	defer wg.Done()

	for w := range pool {
		if worker.status == false {
			fmt.Printf("worker:%d spawning\n", worker.id)
			worker.status = true
		}
		fmt.Printf("worker:%d sleep:%0.1f", worker.id, w)
		time.Sleep(1)
	}
	if worker.status == true {
		fmt.Printf("worker:%d stopping", worker.id)
	}
}

func Pool(max_num int) {
	var wg sync.WaitGroup
	var id int = 1
	pool := make(chan float64, max_num)
	worker := make([]Worker, max_num)

	reader := bufio.NewScanner(os.Stdin)
	reader.Split(bufio.ScanLines)
	for reader.Scan() {
		value, _ := strconv.ParseFloat(reader.Text(), 64)
		pool <- value
		if id <= max_num {
			worker[id-1].id = id
			worker[id-1].status = false
			wg.Add(1)
			go Job(pool, worker[id-1], &wg)
			id++
		}
	}
	close(pool)
	wg.Wait()
}
