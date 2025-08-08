package main

import "fmt"

func updateName(x string) {
	x = "SHREYASH"
}

func updatemenu ( y map[string]float64) {
	y["Coffee"] = 2.99
}

func main() {
	// Group A -> strings , inits , bols , floats , arrays ,structs
	name := "DESAI"

	updateName(name)

	fmt.Println(name)

	// group B -> Slices , Maps , Funcations
	menu := map[string]float64{
		"Pie" : 5.55 , 
		"Ice Cream" : 2.99,	
	}
	updatemenu(menu)
	fmt.Println(menu)
}