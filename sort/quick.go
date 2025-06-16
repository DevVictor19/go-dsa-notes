package sort

/*

# Quick Sort - O(n log n)

It works on the principle of divide and conquer, breaking down the problem into smaller sub-problems.

There are mainly three steps in the algorithm:

1. Choose a Pivot: Select an element from the array as the pivot. The choice of pivot
can vary (e.g., first element, last element, random element, or median).

2. Partition the Array: Rearrange the array around the pivot. After partitioning, all
elements smaller than the pivot will be on its left, and all elements greater than the pivot
will be on its right. The pivot is then in its correct position, and we obtain the index of
the pivot.

3. Recursively Call: Recursively apply the same process to the two partitioned sub-arrays
(left and right of the pivot).

4. Base Case: The recursion stops when there is only one element left in the sub-array, as a
single element is already sorted.

## Complexity Analysis of Quick Sort
- Best Case: (Ω(n log n))
- Average Case (θ(n log n))
- Worst Case: (O(n²)) Occurs when the smallest or largest element is always
chosen as the pivot (e.g., sorted arrays).

## Auxiliary Space
- O(n), due to recursive call stack

*/

func QuickSort(arr []int, low, high int) {
	if low < high {

		// pi is the partition return index of pivot
		pi := partition(arr, low, high)

		// Recursion calls for smaller elements
		// and greater or equals elements
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// Choose the pivot
	pivot := arr[high]

	// Index of smaller element and indicates
	// the right position of pivot found so far
	i := low - 1

	// Traverse arr[low..high] and move all smaller
	// elements on left side. Elements from low to
	// i are smaller after every iteration
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Move pivot after smaller elements and
	// return its position
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}
