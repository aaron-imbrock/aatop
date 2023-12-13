package main

import (
	"fmt"

	"github.com/aaron-imbrock/aatop/read"
)

func main() {
	fmt.Print(read.File("test.txt"))
}
