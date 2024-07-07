package main

import (
	"flag"
	"fmt"
	"goBBP/bbp"
)

var start = flag.Int("start", 0, "The starting digit")
var num = flag.Int("num", 12, "The number of digits")

func main() {
	flag.Parse()
	pi := bbp.New()
	fmt.Println(pi.GetDecimalValues(*start, *num))
}
