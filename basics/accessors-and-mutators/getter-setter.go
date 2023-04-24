package main

import (
	stuff "example/project/mypackage"
	"fmt"
	"log"
)

func main() {
	date := stuff.Date{}
	err := date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetYear(1974)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("1st Day: %d/%d/%d\n", date.Month(), date.Day(), date.Year())
}
