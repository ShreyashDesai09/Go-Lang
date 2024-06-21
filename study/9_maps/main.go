package main

import "fmt"

func main() {

	menu := map[string]float64{
		"Soup":   4.99,
		"pie":    7.99,
		"salad":  6.99,
		"toffee": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//Looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key
	phonebook := map[int]string{
		1: "SD",
		2: "VP",
		3: "MG",
	}

	fmt.Println(phonebook)

	phonebook[2] = "Vishwajeet Patil"
	fmt.Println(phonebook)
}
