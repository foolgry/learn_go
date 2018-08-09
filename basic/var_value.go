package main

import "fmt"

var java, py, golang = "java", "python", 666.66

func main() {
	var i, j, k = 1, 32, 4
	a := 10
	fmt.Println(i, j, k, java, py, golang, a)
	fmt.Printf("golang type:%T\n", golang)
	fmt.Printf("java type:%T", java)
}
