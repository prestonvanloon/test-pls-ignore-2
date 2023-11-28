package main

import (
	"fmt"
	"os"

	"github.com/prestonvanloon/test-pls-ignore/lib/hello"
)

func main() {
	var name string
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Println(hello.Hello(name))
}
