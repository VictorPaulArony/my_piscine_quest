/*package main

import (
	"os"


	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		for _, r := range arg {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}*/
//last param
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	last := args[len(args)-1]
	if len(args) > 0 {
		for _, char := range last {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

/*package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	index := 0
	if index >= 0 && index < len(args) {
		arg := args[index]
		fmt.Println(arg)
	}
}
*/
