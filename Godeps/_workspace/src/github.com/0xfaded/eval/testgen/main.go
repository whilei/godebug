package main

import (
	"fmt"
	"github.com/0xfaded/go-testgen"
	"os"
)

func main() {
	if err := testgen.Generate(&Test{}, os.Stdout); err != nil {
		panic(fmt.Sprintf("Test generation failed %v\n", err))
	}
}
