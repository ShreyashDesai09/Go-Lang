package main

import "fmt"

func main() {
	name := "SD"
	fmt.Println(name)

	fmt.Println("POINTER:",&name)

	m := &name
	fmt.Println("Memory address:",m)
	fmt.Println("Value at memory address:",*m)
}