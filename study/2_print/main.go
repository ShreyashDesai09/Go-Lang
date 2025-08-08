package main

import "fmt"

func main() {
	name := "Shreyash Desai"
	fmt.Println(name) //line printing

	loc := "Kolhapur"
	fmt.Print(loc) //print

	house := "vastushree garden"
	fmt.Printf("\nI live at %s\n", house) //Simple Printf like C Language

	no := 301
	fmt.Printf("%d \n" , no)
	hno := fmt.Sprintf("%d",no)
	fmt.Println("House no", hno)

}
