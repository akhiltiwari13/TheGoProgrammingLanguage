// main packages messes around with the 2 package that we have created earlier...
package main

import (
	"fmt"
	"os"
	"strconv"

	"./packages/tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroK)

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s and %s, %s = %s and %s \n",
			f, tempconv.FToC(f), tempconv.FToK(f), c, tempconv.CToF(c), tempconv.CToK(c))
	}
}
