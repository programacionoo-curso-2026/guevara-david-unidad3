package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	ID     int
	Status string
}

func main() {
	orders := generateOrders(20)

	processOrders(orders)

	fmt.Printf("Numero de Orderenes %d\n", len(orders))
	fmt.Print("Todas las operaciones completadas. Finalizando\n")
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

func processOrders(orders []*Order) {
	for _, order := range orders {
		delay := rand.Intn(500)
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Printf("Procesando orden %d en %dms\n", order.ID, delay)
	}
}
