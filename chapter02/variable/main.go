package main

import "fmt"

//global var

var city string = "Esf"
var area int = 4

func main() {

	//explict
	var home string = "Iran"

	// inferred
	var firstname, lastname = "reza", "saberi"

	//shorthand
	village := "Semirom"
	father := "rahim"

	//const
	const shape = "tall , boy"

	fmt.Println(home, city, area, firstname, lastname, village, father, shape)
}
