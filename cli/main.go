package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/UltiRequiem/roman"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: roman <number>")
		os.Exit(1)
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	numeral, err := roman.ConvertToRoman(number)
	if err != nil {
		panic(err)
	}

	fmt.Println(numeral)
}
