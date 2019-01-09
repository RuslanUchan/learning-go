package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	c1 := <-process(1, 3, &wg)
	c2 := <-process(2, 1, &wg)
	c3 := <-process(3, 5, &wg)

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)

	wg.Wait()
}

func process(id int, delay time.Duration, wg *sync.WaitGroup) <-chan string {
	output := make(chan string)

	go func() {
		// top function declaration using defer
		// the process, in this case done, will be deferred until the
		// function has done executing

		// fmt.Println(*wg)
		// fmt.Println(wg)
		defer wg.Done()

		time.Sleep(time.Second * delay)

		output <- fmt.Sprintf("Process %d done", id)
	}()

	return output
}
