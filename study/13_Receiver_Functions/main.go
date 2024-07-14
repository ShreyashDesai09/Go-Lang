package main

import "fmt"

func main() {
	mybill := newbill("SD's Bill")

	mybill.format()

	fmt.Println(mybill.format())
}