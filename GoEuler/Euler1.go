package main

import (
	"fmt"
	"sync"
)

func main() {
	comm1 := make(chan int, 1)
	comm2 := make(chan int, 1)
	comm3 := make(chan int, 1)

	var totalerWaitGroup sync.WaitGroup
	var goRoutinesDone [3]bool

	goRoutinesDone[0] = false
	goRoutinesDone[1] = false
	goRoutinesDone[2] = false

	totalerWaitGroup.Add(3)

	var tot [3]int
	go func() {
		defer totalerWaitGroup.Done()
		tot[0] = totaler(comm1, &goRoutinesDone[0])
	}()

	go func() {
		defer totalerWaitGroup.Done()
		tot[1] = totaler(comm2, &goRoutinesDone[1])
	}()

	go func() {
		defer totalerWaitGroup.Done()
		tot[2] = totaler(comm3, &goRoutinesDone[2])
	}()

	go subadder(3, 1000, comm1, &goRoutinesDone[0])
	go subadder(5, 1000, comm2, &goRoutinesDone[1])
	go subadder(15, 1000, comm3, &goRoutinesDone[2])

	totalerWaitGroup.Wait()

	total := tot[0] + tot[1] - tot[2]
	fmt.Println(total)
}

func subadder(mod int, n int, t chan int, doneFlag *bool) {
	for i := mod; i < n; i+=mod {
		t <- i
	}
	*doneFlag = true
}

func totaler(t chan int, done *bool) int{
	total := 0;

	for *done != true {
		next := <-t
		total += next
	}

	return total
}
