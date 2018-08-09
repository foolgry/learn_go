package main

import "fmt"

func add2(x, y int) (int) {
	return x + y;
}

func add3(x, y int) (sum int) {
	sum = x + y
	return
}

func main() {
	fmt.Println(add2(2,4))
	fmt.Println(add3(2,4))
}