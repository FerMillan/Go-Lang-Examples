package main

import (
	"awesomeProject/modules/aritmetic"
	"fmt"
)

func main() {
	fmt.Println("hola")
	res := aritmetic.Sum(5, 6)

	fmt.Println(res)
}
