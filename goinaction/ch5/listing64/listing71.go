package main

import (
	"fmt"

	"github.com/mymikasa/mktmp/goinaction/ch5/listing64/entities"
)

func main() {
	a := entities.Admain{
		Rights: 10,
	}

	a.Name = "mikasa"
	a.Email = "mikasa@email.com"
	fmt.Printf("User: %v\n", a)
}
