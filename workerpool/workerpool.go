package main

import(
	"bufio"
	"fmt"
	"sync"
	"os"
	"time"
)
var(
	wrk = []struct{
		MaxThreads int
		Input []string
		WaitDur int
	}{
		{
			2,
			[]string{
				"0.1",
				"0.2",
			},
			1,
		},
	}
)
func StartQueue(queue int) {
	proces := make(chan string, queue)

	var wg sync.WaitGroup
	i := 1

	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		str := scan.Text()
		proces <- str
		if i <= queue {
			wg.Add(1)
			go worker(i, &wg, proces)
			i++
		}
	}
	close(proces)
	wg.Wait()
}

func worker(i int, wg *sync.WaitGroup, proces chan string) {
	defer wg.Done()
	if len(proces) > 0 {
		fmt.Printf("worker: &d spawing\n", i)
		for j := range proces {
			t, _ := time.ParseDuration(j + "s")
			fmt.Printf("worker: %d sleep: %s\n", i, j)
			time.Sleep(t)
		}
		fmt.Printf("worker: %d stopping\n", i)
	}
}

func test(queue int) {
	run := &sync.WaitGroup{}
	run.Add(1)

	go func() {
		defer run.Done()
		StartQueue(queue)
	}

	return run
}

func main() {
	run := test(wrk.MaxThreads)
}
