// echo prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {

	for index, arg := range os.Args[1:] {
		fmt.Print(index, " ")
		fmt.Println(arg)
	}
}
