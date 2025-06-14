package main

import (
	"fmt"

	"github.com/DevVictor19/dsa/search"
	"github.com/DevVictor19/dsa/sort"
)

// "github.com/DevVictor19/dsa/search"

func main() {
	arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

	// Binary Search Examples
	fmt.Println(search.BinarySearch(arr, 2))
	fmt.Println(search.BinarySearchRecursive(arr, 0, len(arr)-1, 23))

	// Linear Search Examples
	fmt.Println(search.LinearSearch(arr, 5))

	// Bubble sort Examples
	arr2 := []int{64, 34, 25, 12, 22, 11, 90}
	sort.BubbleSort(arr2)
	fmt.Println(arr2)
}
