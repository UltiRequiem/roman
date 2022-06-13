package main

import (
	"fmt"
	"github.com/UltiRequiem/roman"
)

func main() {
	roman, error := roman.ConvertToRoman(33)

	if error != nil {
                fmt.Println(error)
                return
	}

	fmt.Println(roman) //=> XXXIII
}
