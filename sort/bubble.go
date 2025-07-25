package sort

/*

# Bubble Sort Algorithm - O(n²)

- Works by repeatedly swapping the adjacent elements if they are in the wrong order.

- This algorithm is not suitable for large data sets as its average and worst-case
time complexity are quite high.


## Complexity Analysis of Bubble Sort
- Time Complexity: O(n²)
- Auxiliary Space: O(1)

# Advantages of Bubble Sort:
- Bubble sort is easy to understand and implement.
- It does not require any additional memory space.
- It is a stable sorting algorithm, meaning that elements with the same key value
maintain their relative order in the sorted output.

# Disadvantages of Bubble Sort:
- Bubble sort has a time complexity of O(n2) which makes it very slow for large data sets.
- Bubble sort has almost no or limited real world applications. It is mostly used in
academics to teach different ways of sorting.

## How it works?

- We sort the array using multiple passes. After the first pass, the maximum element
goes to end (its correct position). Same way, after second pass, the second largest
element goes to second last position and so on.

- In every pass, we process only those elements that have already not moved to correct
position. After k passes, the largest k elements must have been moved to the last k
positions.

- In a pass, we consider remaining elements and compare all adjacent and swap if larger
element is before a smaller element. If we keep doing this, we get the largest
(among the remaining elements) at its correct position.

*/

func BubbleSort(arr []int) {
	length := len(arr)
	swapped := false

	for i := 0; i < length; i++ {
		swapped = false
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				aux := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = aux
				swapped = true
			}
			// fmt.Printf("iteration %d | arr %v\n", i+1, arr)
		}
		if !swapped {
			break
		}
	}
}
