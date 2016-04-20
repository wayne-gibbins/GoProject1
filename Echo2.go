// Echo3 - prints command line args

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

