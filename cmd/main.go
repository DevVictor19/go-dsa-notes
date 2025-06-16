package main

import (
	"fmt"

	"github.com/DevVictor19/dsa/search"
	"github.com/DevVictor19/dsa/sort"
)

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

	// Merge sort Examples
	arr3 := []int{12, 11, 13, 5, 6, 7}
	sort.MergeSort(arr3, 0, len(arr3)-1)
	fmt.Println(arr3)

	// Insertion sort Examples
	arr4 := []int{12, 11, 13, 5, 6}
	sort.InsertionSort(arr4)
	fmt.Println(arr4)
}
