// Echo5 prints its command-line arguments and index per line
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Printf("%d %s\n", index, arg)
	}
}
