package main

import (
	"runtime"
	"fmt"
)

func main() {
	switch os:= runtime.GOOS; os {
	case "drawin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("liunx")
	default:
		fmt.Printf("%s", os)
	}
}
