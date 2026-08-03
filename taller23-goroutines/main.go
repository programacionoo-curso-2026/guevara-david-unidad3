package main

import (
	/*"fmt"*/
	/*"math/rand"*/
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
