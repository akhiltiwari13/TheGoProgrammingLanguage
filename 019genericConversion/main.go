// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./packages/units"
)

func main() {
	// check if the conversion number are passed as arguments are not.
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := units.Fahrenheit(t)
			c := units.Celsius(t)
			m := units.Meters(t)
			ft := units.Feet(t)
			lb := units.Pounds(t)
			kg := units.Kilograms(t)

			fmt.Printf("%s = %s and %s, %s = %s and %s \n",
				f, units.FToC(f), units.FToK(f), c, units.CToF(c), units.CToK(c))

			fmt.Printf("%s = %s , %s = %s \n",
				m, units.MtToFt(m), ft, units.FtToMts(ft))

			fmt.Printf("%s = %s , %s = %s \n",
				lb, units.PdToKg(lb), kg, units.KgToPd(kg))

		}
	} else { // else section needs to be completed, it deals with taking input from the stdin contrary to arguments to main.
		// take the inputs from the stdin.
		// var strInputs []string
		// input := bufio.NewScanner(os.Stdin)
		// for input.Scan() {
			// strInputs[]
		}

	}

}
