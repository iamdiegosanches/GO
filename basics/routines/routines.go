package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

func printTo15() {
	for i := 1; i <= 15; i++ {
		pl("Fun 1:", i)
	}
}

func printTo10() {
	for i := 1; i <= 10; i++ {
		pl("Fun 2:", i)
	}
}

func main() {
	go printTo15()
	go printTo10()
	time.Sleep(2 * time.Second)
}
