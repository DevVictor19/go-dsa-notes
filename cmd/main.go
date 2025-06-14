package main

import (
	"fmt"

	"github.com/DevVictor19/dsa/search"
)

func main() {
	// Binary Search Examples
	arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	fmt.Println(search.BinarySearch(arr, 2))
	fmt.Println(search.BinarySearchRecursive(arr, 0, len(arr)-1, 23))
}
