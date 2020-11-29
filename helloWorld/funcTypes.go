package main

import "fmt"

type aNewFunction func(s string)

func invokeClose(c aNewFunction) {
	c("Hello world")
}

func main() {
	// A function can also be stored in a variable
	invokeClose(func(s string) {
		fmt.Println(s)
	})
}
