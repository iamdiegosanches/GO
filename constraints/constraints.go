package main

import (
	"fmt"
)

var pl = fmt.Println

type MyContraint interface {
	int | float64
}

func getSumGen[T MyContraint](x T, y T) T {
	return x + y
}

func main() {
	pl("5 + 4 =", getSumGen(5,4))
	pl("5.6 + 4.6 =", getSumGen(5.6,4.6))
}
