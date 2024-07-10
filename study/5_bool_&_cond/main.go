package main

import "fmt"

func main() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is not less than 40")
	}

	names := []string{"SD", "VP", "MG", "VK", "SZ"}

	for index, value := range names {
		if index == 1 {
 			fmt.Println("\nContinue at pos", index , "\n")
			continue
		}
		if index > 3 {
			fmt.Println("Breaking at pos", index)
			fmt.Printf("Value at beaking index %v is %v", index, value)
			break
		}
		fmt.Printf("The value at pos %v is %v\n", index, value)
	}
}
