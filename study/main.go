package main

import (
	"fmt"
)

func main () {
	
	ANS := "SD" 
	var SD string
	fmt.Println("ENTER VALUE:")
	fmt.Scan(&SD)
	fmt.Println(SD)

	if (SD != ANS) {	
		fmt.Println("WRONG ENTER")
	} else {
		fmt.Println("RIGHT ENTER")
	}

	var SW uint32 
	a := 10
	b := 5

	fmt.Printf("CHOOSE:-\n1 ADD \n2 SUB\n3 DIV \n4 MUL\n")
	fmt.Scan( &SW)
	fmt.Println("ENTERED ANS " , SW)

	switch SW {
	case 1: 
	fmt.Println("ADDITION ",a+b)
	break;

	case 2 :
		fmt.Println("SUBSTRACTION" , a-b)

	default : 
	fmt.Println("NO SUCH VALUE")
	break;
		
	}
	
}