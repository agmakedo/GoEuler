package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

const TARGET = 600851475143

func main() {
	fmt.Println(LargestPrimeFactor(TARGET))
}

func IsPrime(input int) bool {
	if input%2 == 0 {
		return false
	}
	for i := 3; i < input/2; i += 2 {
		if input%i == 0 {
			return false
		}
	}
	return true
}

func LargestPrimeFactor(input int) int {
	cores := runtime.NumCPU()

	if cores == 0 {
		cores = 1
	}

	result := 0
	if input%2 == 0 {
		result = 2
	}
	searchlim := int(math.Sqrt(float64(input + 1)))
	searchinc := searchlim / cores

	SearchRange := func(section int, increment int) (int, int) {
		searchmin := section*increment + 1
		searchmax := (section+1)*increment + 1
		return searchmin, searchmax
	}

	FindFactor := func(bot int, top int) int {
		res := 1
		for i := bot; i < top; i++ {
			if input%i == 0 && IsPrime(i) {
				res = i
			}
		}
		return res
	}

	var wg sync.WaitGroup

	wg.Add(cores)

	f := func(bot int, top int) {
		defer wg.Done()
		x := FindFactor(bot, top)
		if x > result {
			result = x
		}
	}

	for i := 0; i < cores-1; i++ {
		go f(SearchRange(i, searchinc))
	}
	go f(searchlim-searchinc-1, searchlim)

	wg.Wait()

	return result
}
