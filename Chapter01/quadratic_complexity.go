package main

import (

	"fmt"
)

func main() {

	var k, l int 

	for k = 1; k <= 10; k++ {

		fmt.Println("multiplication Table", k)

		for l = 1; l <= 10; l++ {

			var x = k * l 
			fmt.Println(x)
		}
	}
}