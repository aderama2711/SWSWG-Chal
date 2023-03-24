package main

import (
	"fmt"
	"sync"
)

type inter interface {
	print(i int)
}

type data struct {
	isinya []string
}

func (d data) print(i int) {
	fmt.Println(d.isinya, i)
}

func main() {
	// Without mutex
	var wg sync.WaitGroup

	data_coba := []string{"coba1", "coba2", "coba3"}
	data_bisa := []string{"bisa1", "bisa2", "bisa3"}

	var c1 inter = data{isinya: data_coba}
	var c2 inter = data{isinya: data_bisa}

	fmt.Println("Without Mutex")

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go goroutine(c1, i, &wg)
		wg.Add(1)
		go goroutine(c2, i, &wg)
	}

	wg.Wait()

	// With mutex

	var mtx sync.Mutex

	fmt.Println("With Mutex")

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go goroutinewithmutex(c1, i, &wg, &mtx)
		wg.Add(1)
		go goroutinewithmutex(c2, i, &wg, &mtx)
	}

	wg.Wait()

}

func goroutine(c inter, i int, wg *sync.WaitGroup) {
	c.print(i)
	wg.Done()
}

func goroutinewithmutex(c inter, i int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	mtx.Lock()
	c.print(i)
	mtx.Unlock()
	wg.Done()
}
