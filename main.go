package main

import (
	"flag"
	"fmt"
	"goBBP/bbp"
	"strconv"
)

var startArg = flag.String("start", "0", "The starting digit")
var numArg = flag.String("num", "12", "The number of digits")

func main() {
	start, err := strconv.Atoi(*startArg)
	if err != nil {
		fmt.Println("Error converting startArg", err)
		return
	}
	num, err := strconv.Atoi(*numArg)
	if err != nil {
		fmt.Println("Error converting startArg", err)
		return
	}
	pi := bbp.New()
	fmt.Println(pi.GetDecimalValues(start, num))
}
