// "Strange" array manipulation

package main

import "fmt"

func main() {
   var myArray [5]int

   // Set values in the array
   myArray[0] = 10
   myArray[1] = 20
   myArray[2] = 30
   myArray[3] = 40
   myArray[4] = 50

   // Access a value in the array
   fmt.Println("The third element in the array is:", myArray[2])
   
   // Declare and initialize an array with shorthand syntax
   anotherArray := [3]string{"apple", "banana", "cherry"}

   // Print the entire array
   fmt.Println("The second array is:", anotherArray)

   // Loop through the array and print each value
   for i := 0; i < len(myArray); i++ {
       fmt.Println("Element", i, "is:", myArray[i])
   }	
  
  // ------------ Exercises -----------------
  a := [...]int{0, 1, 2, 3}
	x := a[:1]
	y := a[2:]
	x = append(x, y...)
	x = append(x, y...)
	fmt.Println(a, x)
  
  // output: [0 2 3 3] [0 2 3 3 3]
  // Editing a slice will edit the content of the original array
  
   var y = []string{"A", "B", "C", "D"}
   var x = y[:3]
   fmt.Printf("%v", x)
   for i, s := range x {
	print(i, s, ",")
	x = append(x, "Z")
	x[i+1] = "Z"
   }
  // output: 0A,1Z,2C,
  
   var x = []string{"A", "B", "C"}

   for i, s := range x {
	print(i, s, " ")
	x = append(x, "Z")
	x[i+1] = "Z"
   }
   // output: 0A,1B,2C,
  
   var x = []string{"A", "B", "C"}

   for i, s := range x {
	print(i, s, ",")
	x[i+1] = "M"
	x = append(x, "Z")
	x[i+1] = "Z"
   }
  
  // output: 0A,1M,2C,
  
  /*
  Key points:

    Ranging over a container iterates the elements of a copy of the container. The evaluation of the copy happens before executing the loop, 
    so the length and capacity of the copy are never changed.

    A slice references its elements on a backing array. So a copy of a slice shares the same elements (and the backing array) with the slice.

    The assignment x[i+1] = "M" in the first loop step modifies the second element of the initial slice x and the copy of the initial slice x.
    If the free element slots of the first argument slice of an append call are insufficient to hold all the appended elements, 
    the append call will create a new backing array to hold all the elements of the first argument slice and the appended elements. 
    So, at the end of the first loop step, the backing array of the slice x is changed. 
    However, the change doesn't affect the slice copy used in the element iteration. 
    All subsequent element modifications apply to the new backing array, so they have no effects on the copy used in the element iteration.
  
  */
}
