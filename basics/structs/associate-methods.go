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

func (tsp Tsp) toMLs() ML {
	return ML(tsp * 4.92)
}


func (tbs TBs) toMLs() ML {
	return ML(tbs * 14.79)
}

func main() {
	tsp1 := Tsp(3)
	fmt.Printf("%.2f tsp = %.2f ml\n", tsp1, tspToML(3))
	
}
