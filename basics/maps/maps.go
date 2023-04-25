package main

import "fmt"

func main() {
  // Array for test
	nums := [...]int{1, 2, 3, 3}
  
	m := make(map[int]int)
	for _, n := range nums {
		if _, ok := m[n]; ok { // Verify if the key aready existis
			return true
		} else {
			m[n] = n
		}
	}
	return false
}
