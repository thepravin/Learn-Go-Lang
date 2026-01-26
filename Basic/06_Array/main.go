/**
- length of array declared and fixed at the time of decleration.
- when you declare an array or a slice, the elements are initialized to their  default value.
	int,float = 0
	string 	  = ""
	boolean   = false
	pointers  = nil
*/

package main

import "fmt"

func main() {
	var numbers [3]int

	numbers[0] = 1
	numbers[1] = 2

	fmt.Println("Numbers : ", numbers) // 3rd element is 'zero'

	fmt.Println("Number at index 1 : ", numbers[1])

	//--------------------------------------------------------------------

	fruits := [5]string{"apple", "orange"}

	fmt.Printf("Array: %q\n", fruits) // Will print: ["apple" "orange" "" ""  ""] (note the 3 empty spots)
	// %q for the quotes.

	//--------------------------------------------------------------------

	fmt.Println("Length of numbers : ", len(numbers))
	fmt.Println("Length of fruits : ", len(fruits))
}
