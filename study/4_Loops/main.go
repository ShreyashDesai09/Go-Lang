package main

import "fmt"

func main() {
	//long way
	x := 0
	for x < 5 {
		fmt.Println("value of x is:",x)					
		x++
	}

	//short way
	for i := 0;i<5;i++ {
		fmt.Println("value of i:",i)
	}
	//value
	names := []string {"SD","VP","MG"}

	for i:= 0;i<len(names);i++ {
		fmt.Println("Name is",names[i])
	}

	//with index and value
	for index,value := range names {
		fmt.Printf("Position at index %v is %v\n",index,value)
	}

	fmt.Println(names)
}