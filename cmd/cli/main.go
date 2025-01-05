package main

import (
	"fmt"
	"os"

	"github.com/boolka/gohello/v2/pkg/hello"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(hello.Hello())
		return
	}

	for i, arg := range os.Args {
		if i == 0 {
			continue
		}

		switch arg {
		case "-h", "--help":
			fmt.Println("github.com/boolka/gohello")
		case "-p", "--proverb":
			fmt.Println(hello.Proverb())
		}
	}
}
