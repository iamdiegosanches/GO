package main

import (
	"fmt"
)

var pl = fmt.Println

type Tsp float64
type TBs float64
type ML float64

func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

func tBToML(tsp TBs) ML {
	return ML(tsp * 14.79)
}

func main() {
	ml1 := ML(Tsp(3) * 4.92)
	fmt.Printf("3 tsps = %.2f ML\n", ml1)
	ml2 := ML(TBs(3) * 14.79)
	fmt.Printf("3 TBs = %.2f ML\n", ml2)
	pl("2 tsp + 4 tsp = ", Tsp(2) + Tsp(4))
	pl("2 tsp > 4 tsp = ", Tsp(2) > Tsp(4))
	fmt.Printf("3 tsp = %.2f ml\n", tspToML(3))
	fmt.Printf("3 tsp = %.2f ml\n", tBToML(3))
}
