package main

import "fmt"

var v float64
var k string

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newbill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}
	return b
}

// format bill
func (b bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0

	//list items
	for k, v= range b.items {
		fs += fmt.Sprintf("%v ...$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%v ...$%0.2f" , "Total:" , total)

	return fs
}