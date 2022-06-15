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

	number, errorParsingInput := strconv.Atoi(os.Args[1])

	if errorParsingInput != nil {
		fmt.Println(errorParsingInput)
		os.Exit(1)
	}

	numeral, errrorConvertingInput := roman.ConvertToRoman(uint16(number))

	if errrorConvertingInput != nil {
		fmt.Println(errorParsingInput)
		os.Exit(1)
	}

	fmt.Println(numeral)
}
