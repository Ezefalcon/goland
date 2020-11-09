package main

import (
	"errors"
	"fmt"
)

func main() {
	sum, err := sum(1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}
}

func sum(a int, b int) (int, error) {

	if a < b {
		return 0, errors.New("el primer valor es mayor al segundo")
	}
	return a+b, nil
}


