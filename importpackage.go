package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		n := rand.Intn(100)
		fmt.Println(n)
	}
}
