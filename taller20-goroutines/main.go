package main

import (
	"fmt"
)

func main() {
	ShowGoroutines(1)
}

func ShowGoroutines(id int) {
	fmt.Printf("Goroutines #%d\n", id)
}
