package main

import (
	"fmt"
	"strconv"

	c "github.com/damiisdandy/zero-to-billion-go/converter"
)

func main() {
	var value string
	fmt.Print("Enter a number: ")
	fmt.Scan(&value)
	if num, err := strconv.Atoi(value); err == nil {
		if num < 0 || num > c.MAX_UPPER_LIMIT {
			fmt.Println("Number is out of range")
			return
		}
		fmt.Printf("converted to words is: %q\n", c.Converter(num))
	} else {
		fmt.Println("The input is not an integer number")
	}
}
