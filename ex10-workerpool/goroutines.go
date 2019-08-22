package goroutines

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

func job(personal Worker, pool chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range pool {
		if personal.status == false {
			fmt.Printf("worker:%d spawning\n", personal.id)
			personal.status = true
		}
		fmt.Printf("worker:%d sleep:%.1f\n", personal.id, job)
		time.Sleep(1 * time.Second)
	}
	if personal.status == true {
		fmt.Printf("worker:%d stopping\n", personal.id)
	}
}

func Run(poolSize int) {
	var wg sync.WaitGroup
	var id int = 1
	personal := make([]Worker, poolSize)
	pool := make(chan float64, poolSize)

	reader := bufio.NewScanner(os.Stdin)
	reader.Split(bufio.ScanLines)
	for reader.Scan() {
		number, _ := strconv.ParseFloat(reader.Text(), 64) //// shift to top
		pool <- number
		if id <= poolSize {
			personal[id-1].id = id
			personal[id-1].status = false
			wg.Add(1)
			go job(personal[id-1], pool, &wg)
			id++
		}
	}
	close(pool)
	wg.Wait()
}
