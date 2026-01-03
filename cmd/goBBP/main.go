package main

import (
	"flag"
	"fmt"
	"goBBP/internal/bbp"
)

var (
	pStart = flag.Int("pStart", 0, "The starting digit")
	pNum   = flag.Int("pNum", 50, "The number of digits")
	pHex   = flag.Bool("hex", false, "Output hex instead of decimal")
	pLower = flag.Bool("lower", false, "Use lowercase for hex output")
)

func main() {
	flag.Parse()
	pi := bbp.New()

	if *pHex {
		if *pLower {
			fmt.Println(pi.GetHexLowerString(*pStart, *pNum))
		} else {
			fmt.Println(pi.GetHexString(*pStart, *pNum))
		}
	} else {
		fmt.Println(pi.GetDecimalString(*pStart, *pNum))
	}
}
