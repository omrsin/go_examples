// Echo4 just prints its command-line arguments and the command itself
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
