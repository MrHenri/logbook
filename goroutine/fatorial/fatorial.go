package fatorial

import (
	"log"
	"time"
)

func RecursiveFat(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * RecursiveFat(num-1)
}

func IterativeFat(num int) int {
	r := 1
	for i := 1; i <= num; i++ {
		r *= i
	}
	return r
}

func IterativeFatConcurrency(startNum int, endNum int, resultChan chan int) {
	r := 1
	for i := startNum; i <= endNum; i++ {
		r *= i
	}
	resultChan <- r
}

func ConcurrencyFat(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	rChan := make(chan int)
	mid := num / 2

	go IterativeFatConcurrency(1, mid, rChan)
	go IterativeFatConcurrency(mid+1, num, rChan)

	y, z := <-rChan, <-rChan
	return y * z

}

func ConcurrencyFatBySort(num int) int {
	if num == 0 || num == 1 {
		return 1
	}

	var l []int
	var r []int
	rChan := make(chan int)

	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			l = append(l, i)
		} else {
			r = append(r, i)
		}
	}

	go MultiplyArray(r, rChan)
	go MultiplyArray(l, rChan)
	x, y := <-rChan, <-rChan

	return x * y

}

func MultiplyArray(nums []int, resultChan chan int) {
	r := 1
	for _, num := range nums {
		r *= num
	}
	resultChan <- r
}

func fatSpeedTest(i int) {
	{
		start := time.Now()
		r := RecursiveFat(i)
		elapsed := time.Since(start)
		log.Printf("RecursiveFat result: %d\n", r)
		log.Printf("RecursiveFat took %d\n\n", elapsed.Nanoseconds())
	}
	{
		start := time.Now()
		r := IterativeFat(i)
		elapsed := time.Since(start)
		log.Printf("IterativeFat result: %d\n", r)
		log.Printf("IterativeFat took %d\n\n", elapsed.Nanoseconds())
	}
	{
		start := time.Now()
		r := ConcurrencyFatBySort(i)
		elapsed := time.Since(start)
		log.Printf("ConcurrencyFatBySort result: %d\n", r)
		log.Printf("ConcurrencyFatBySort took %d\n\n", elapsed.Nanoseconds())
	}
	{
		start := time.Now()
		r := ConcurrencyFat(i)
		elapsed := time.Since(start)
		log.Printf("ConcurrencyFat result: %d\n", r)
		log.Printf("ConcurrencyFat took %d", elapsed.Nanoseconds())
	}
}
