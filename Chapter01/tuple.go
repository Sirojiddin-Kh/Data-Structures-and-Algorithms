package main

import (

	"fmt"
)

func powerSeries(num int) (int, int) {


	return num * num, num * num * num
}

func main() {

	var square, cube int 

	square, cube = powerSeries(3)

	fmt.Println("Square ", square, "Cube ", cube )
}