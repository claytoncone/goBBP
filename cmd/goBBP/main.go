package main

import (
	"flag"
	"fmt"
	"goBBP/internal"
)

var pStart = flag.Int("pStart", 1000000, "The starting digit")
var pNum = flag.Int("pNum", 30, "The number of digits")

func main() {
	flag.Parse()
	pi := bbp.New()
	fmt.Println(pi.GetDecimalValues(*pStart, *pNum))
}
