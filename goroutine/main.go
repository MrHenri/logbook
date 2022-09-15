package main

import (
	"log"
	"time"

	"github.com/MrHenri/logbook/goroutine/fatorial"
)

var i int = 20

func main() {
	{
		start := time.Now()
		r := fatorial.RecursiveFat(i)
		elapsed := time.Since(start)
		log.Printf("RecursiveFat result: %d\n", r)
		log.Printf("RecursiveFat took %d\n\n", elapsed.Nanoseconds())
	}
	{
		start := time.Now()
		r := fatorial.IterativeFat(i)
		elapsed := time.Since(start)
		log.Printf("IterativeFat result: %d\n", r)
		log.Printf("IterativeFat took %d\n\n", elapsed.Nanoseconds())
	}
	{
		start := time.Now()
		r := fatorial.ConcurrencyFatBySort(i)
		elapsed := time.Since(start)
		log.Printf("ConcurrencyFatBySort result: %d\n", r)
		log.Printf("ConcurrencyFatBySort took %d\n\n", elapsed.Nanoseconds())
	}
	{
		start := time.Now()
		r := fatorial.ConcurrencyFat(i)
		elapsed := time.Since(start)
		log.Printf("ConcurrencyFat result: %d\n", r)
		log.Printf("ConcurrencyFat took %d", elapsed.Nanoseconds())
	}
}
