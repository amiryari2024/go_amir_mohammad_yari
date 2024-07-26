package main

import (
	"fmt"
)

func isComposite(num int) bool {
	if num < 4 {
		return false
	}
	count := 0
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			count++
		}
	}
	return count > 2
}

func generateComposites(start, end int, ch chan<- int) {
	for i := start; i <= end; i++ {
		if isComposite(i) {
			ch <- i
		}
	}
	close(ch)
}

func processComposites(ch <-chan int, result chan<- string) {
	for composite := range ch {
		square := composite * composite
		if square%2 == 0 {
			result <- "Pong"
		} else {
			result <- "Ping"
		}
	}
	close(result)
}

func main() {
	compositesChan := make(chan int)
	resultsChan := make(chan string)

	go generateComposites(1, 100, compositesChan)
	go processComposites(compositesChan, resultsChan)

	for result := range resultsChan {
		fmt.Println(result)
	}
}
