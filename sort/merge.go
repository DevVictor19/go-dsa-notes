package sort

/*

# Merge Sort  O(n log n)

- It follows the divide-and-conquer approach
- It works by recursively sorting the tow halves, recursively sorting the two halves and finally margin
them back together to obtain the sorted array

## How does Marge Sort Work?

1. Divide:  Divide the list or array recursively into two halves until it can no more be divided.
2. Conquer:  Each sub-array is sorted individually using the merge sort algorithm.
3. Merge:  The sorted sub-arrays are merged back together in sorted order. The process continues
until all elements from both sub-arrays have been merged.

## Time Complexity:
- Best Case: O(n log n)
- Average Case: O(n log n)
- Worst Case: O(n log n)

## Auxiliary Space:
- O(n), Additional space is required for the temporary array used during merging.

*/

func MergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2
	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	n1 := mid - left + 1
	n2 := right - mid

	// Create temp vectors
	L := make([]int, n1)
	R := make([]int, n2)

	// Copy data to temp vectors L[] and R[]
	for i := 0; i < n1; i++ {
		L[i] = arr[left+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[mid+1+j]
	}

	i := 0
	j := 0
	k := left

	// Merge the temp vectors back
	// into arr[left..right]
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	// Copy the remaining elements of L[],
	// if there are any
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	// Copy the remaining elements of R[],
	// if there are any
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}
