package main

import (
	"flag"
	"fmt"
	"goBBP/bbp"
)

var pStart = flag.Int("pStart", 0, "The starting digit")
var pNum = flag.Int("pNum", 50, "The number of digits")

func main() {
	flag.Parse()
	pi := bbp.New()
	fmt.Println(pi.GetDecimalValues(*pStart, *pNum))
}
