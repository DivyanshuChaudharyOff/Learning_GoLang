package main

import "fmt"

func main() {
	var A int = 32768
	A = A + 1
	fmt.Println("A = ", A)
	fmt.Println(" A  store integar value")

	var B float32 = 123.4
	fmt.Println("B = ", B)
	fmt.Println(" B  store float value with 32 bit")

	var C float64 = 123.4
	fmt.Println("C = ", C)
	fmt.Println(" C  store float value with 64 bit")
}
