package main

import (
	"fmt"
)

func main() {

	//Arrays
	var age = [2]int{1, 2}

	name := [2]string{"Shreyash", "SD"}

	fmt.Println(age , len(age))
	fmt.Println(name,len(name))

	//Slices
	var scores= []int{100,200,222}
	scores[2]=300
	scores = append(scores,85)
	fmt.Println(scores,len(scores))
	scores = append(scores, 500)
	fmt.Println(scores,len(scores))

	//slice ranges
	rangeOne := scores[1:3]
	rangeTwo := scores[2:]
	rangeThree := scores[:4]
	fmt.Println(rangeOne , rangeTwo , rangeThree)
}