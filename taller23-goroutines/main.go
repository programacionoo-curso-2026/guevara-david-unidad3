package main

import (
	/*"fmt"*/
	/*"math/rand"*/
	"fmt"
	"sync"
	/*"time"*/)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

var (
	totalUpdates int
	updateMutex  sync.Mutex
)

func main() {
	var wg sync.WaitGroup
	orders := generateOrders(20)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, order := range orders {
				updateOrderStatus(order)
			}
		}()
	}
	wg.Wait()
	fmt.Print("Todas las operaciones completadas. Exiting\n")

}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID: i + 1, Status: "pending",
		}
	}
	return orders
}
