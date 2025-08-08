package main

import (
	"fmt" 
	"math"
)



func sayGreeting(n string) {
	fmt.Printf("Good Morning %v \n",n)
}

func sayBye(n string) {
	fmt.Printf("Good bye %v \n",n)
}

func cycleNames(n []string , f func(string)) {
	for _ ,v := range n {
		f(v)
	}
}

func circleArea(r float64) float64{
	return math.Pi*r*r
}

func main() {
	sayGreeting("Shreyash")
	sayGreeting("SD")
	sayBye("Desai")
	cycleNames([]string{"SD","VP","MG"},sayGreeting)
	cycleNames([]string{"SD","VP","MG"},sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(20)

	fmt.Println(a1,"\n",a2)
}