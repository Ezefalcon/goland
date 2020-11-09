package main

import "fmt"

func main() {
	var arr [2] int
	arr[0] = 1
	arr[1] = 2

	//printArr(arr)

	//var l [] int
	l := make([]int, 10)
	l = append (l, 10)
	fmt.Printf("%p\n", l)
	l = append (l, 100)
	fmt.Printf("%p\n", l)
	l = append (l, 1000)
	fmt.Printf("%p\n", l)

	printArr(l)
	m := make (map [int]string)
	m[0] = "a"
	m[1] = "b"


}

func printArr(arr []int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
