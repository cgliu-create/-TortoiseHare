package main

import (
	"fmt"

	"github.com/cgliu-create/TortoiseHare/algo"
)

func main() {
	fmt.Println("hello")
	ex := []int{4, 3, 2, 7, 5, 8, 3, 6, 1}
	fmt.Println(algo.FindDuplicate(ex))
}
