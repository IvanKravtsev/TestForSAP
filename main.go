package main

import (
	"TestForSAP/mathtest"
	"fmt"
)

func main() {
	var arr mathtest.Arr1000
	arr.Generate()
	fmt.Println(arr.Min())
	fmt.Println(arr.Max())
	fmt.Println(arr.Popular())
	fmt.Println(arr.Unpopular())
	fmt.Println(arr.Mean())
}
