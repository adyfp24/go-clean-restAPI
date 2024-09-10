package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "ady firdaus"
	fmt.Println("Hello, welcome" + name)

	for i := 1; i <= 5; i++ {
		fmt.Println("loop ke : " + strconv.Itoa(i))
	}

}
