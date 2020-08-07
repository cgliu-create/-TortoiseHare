package main

import (
	"fmt"

	"github.com/cgliu-create/TortoiseHare/algo"
	"github.com/cgliu-create/TortoiseHare/duplicateobj"
)

//Point - test struct
type Point struct {
	X, Y int
}

func main() {
	fmt.Println("hello")
	ex := []int{4, 3, 2, 7, 5, 8, 3, 6, 1}
	fmt.Println(algo.FindDuplicate(ex))
	var ex2 = []interface{}{Point{1, 2}, Point{2, 5}, Point{3, 3}, Point{1, 2}, Point{5, 6}, Point{2, 5}}
	fields := []string{"X", "Y"}
	fmt.Println(duplicateobj.FindDuplicateObj(fields, ex2))
}
