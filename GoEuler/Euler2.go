package main

import (
	"fmt"
	"sync"
)

const FIB_MAX = 4000000

func main() {
	total := 0
	var totalerWaitGroup sync.WaitGroup

	lastFibo := 0
	currentFibo := 1

	for lastFibo < FIB_MAX {
		totalerWaitGroup.Add(1)
		temp := lastFibo
		lastFibo = currentFibo
		currentFibo = nextFibo(temp, currentFibo)

		f := func(i int) {
			defer totalerWaitGroup.Done()
			AddIfEven(i, &total)
		}

		go f(currentFibo)
	}

	totalerWaitGroup.Wait()

	fmt.Println(total)
}

func AddIfEven(value int, t *int){
	if value % 2 == 0 {
		*t += value
	}
}

func nextFibo(last int, current int) int {
	return last + current
}

