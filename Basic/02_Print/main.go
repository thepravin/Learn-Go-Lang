package main

import "fmt"

type Product struct {
	ID       int
	Name     string
	IsActive bool
}

func main() {
	age := 22
	name := "Pravin"

	fmt.Println("Age : ", age, "Name : ", name)
	fmt.Println("Next line text")

	fmt.Printf("My age is %d\n", age)
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("May age is %d and Name is %s\n", age, name)

	// If \n not added then cursor not shift to next line it remains same

	fmt.Printf("Type of name is : %T\n", name)

	item := Product{
		ID:       101,
		Name:     "Wireless Mouse",
		IsActive: true,
	}

	fmt.Println("--- Using %v (Values Only) ---")
	fmt.Printf("%v\n", item)

	fmt.Println("\n--- Using %+v (Fields + Values) ---")
	fmt.Printf("%+v\n", item)

}

/*

Verb					Description						Example
%t,Boolean 			(true or false)						true
%p,Pointer address 	(base 16 notation with 0x)		0xc0000b4008

*/
