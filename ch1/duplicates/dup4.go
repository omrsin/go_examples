/* Exercise: Dup4 prints the text of each line that appears
   more than once in the input, preceeded by its count and
   a list of file names on which the duplicated word occurs.
   It reads from standard in or from a list of files.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

type dupInfo struct {
     count int
     files []string
}

func main() {
	counts := make(map[string]*dupInfo)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, d := range counts {
		if d.count > 1 {
			fmt.Printf("%d\t%s\t%v\n", d.count, line, d.files)
		}
	}
}

func countLines(f *os.File, counts map[string]*dupInfo) {
	input := bufio.NewScanner(f)
	for input.Scan(){
		var d *dupInfo
		if _, ok := counts[input.Text()]; !ok {
		   counts[input.Text()] = &dupInfo {
		   	
		   }
		}
		d = counts[input.Text()]
		d.count++
		d.files = append(d.files, f.Name())
	}
}
