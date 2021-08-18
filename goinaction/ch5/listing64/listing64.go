package main

import (
	"fmt"

	"github.com/mymikasa/mktmp/goinaction/ch5/listing64/counters"
)

func main() {
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)
}
