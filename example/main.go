package main

import (
	"fmt"
	"github.com/UltiRequiem/roman"
)

func main() {
	romanNumeral, error := roman.ConvertToRoman(33)

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(romanNumeral) //=> XXXIII

	fmt.Println(roman.ParseRoman("XXXIII")) //=> 33
}
