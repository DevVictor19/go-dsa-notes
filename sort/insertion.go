package sort

/*

# Insertion Sort - O(n²)

- We start with the second element of the array as the first element is assumed to be sorted.

- Compare the second element with the first element if the second element is smaller then swap them.

- Move to the third element, compare it with the first two elements, and put it in its correct
positionRepeat until the entire array is sorted.

## Time Complexity
- Best case: O(n)
- Average case: O(n²)
- Worst case: O(n²)

## Space Complexity
- Auxiliary Space: O(1), Insertion sort requires O(1) additional space,
making it a space-efficient sorting algorithm.
*/

func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		/* Move elements of arr[0..i-1], that are
		   greater than key, to one position ahead
		   of their current position */
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}
