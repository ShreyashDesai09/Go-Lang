package main

import "fmt"

type deck []string

func newDeck() deck {

}

func (d deck) print(){
	for i , card := range d {
		fmt.Println(i,card)
	}
}