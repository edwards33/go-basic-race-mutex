package main

import (
	"fmt"
	"sync"
        "runtime"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var counter int

func main() {
        fmt.Println("Num of CPUs:", runtime.NumCPU())
	wg.Add(2)
	go increment("Proc01:")
	go increment(" Proc02:")
	wg.Wait()
        fmt.Println("Final Count:", counter)
}

func increment(s string) {
	for i := 0; i < 50; i++ {
                mutex.Lock()
                counter++
		fmt.Println(s, i, "Count:", counter)
                mutex.Unlock()
	}
	wg.Done()
}

// go run -race main.go
