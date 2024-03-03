package main

import (
	"fmt"

	"github.com/MeSabya/toolkit"
)

func main() {
	var tools toolkit.Tools
	s, _ := tools.RandomString(5)

	fmt.Println("Random string is", s)
}
