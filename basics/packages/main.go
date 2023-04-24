package main

import (
	"fmt"
	stuff "example/project/mypackage"
)

func main() {
	fmt.Println("Hello", stuff.Name)
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Println(strArr)
}
