package search

/*

# Binary Search Algorithm - O(log n)

Is a search algorithm used in a SORTED array by repeatedly dividing the search interval
in half.

## Conditions to apply
- The data structure must be sorted.
- Access to any element of the data structure should take constant time - O(1)

## Time Complexity
- Best case: O(1)
- Average case: O(log n)
- Worst case: O(log n)

## Auxiliary Space
- O(1)
- if recursive call stack considered O(log n)

*/

/*

Below is the step-by-step algorithm for Binary Search:

- Divide the search space into two halves by finding the middle index "mid".

- Compare the middle element of the search space with the key.

- If the key is found at middle element, the process is terminated.

- If the key is not found at middle element, choose which half will be used as the next search space.
  - If the key is smaller than the middle element, then the left side is used for next search.
  - If the key is larger than the middle element, then the right side is used for next search.

- This process is continued until the key is found or the total search space is exhausted.

*/

func BinarySearch(arr []int, target int) int {
	// iteration := 0
	low := 0
	high := len(arr) - 1

	for low <= high {
		// iteration++
		mid := ((high - low) / 2) + low
		key := arr[mid]

		// fmt.Printf("---------------- iteration: %d\n", iteration)
		// fmt.Println("key:", key)
		// fmt.Println("low:", low)
		// fmt.Println("mid:", mid)
		// fmt.Println("high:", high)

		if key == target {
			//fmt.Printf("Found key: %d - iteration: %d - index: %d\n", key, iteration, mid)
			return mid
		}
		if key > target {
			high = mid - 1
		}
		if key < target {
			low = mid + 1
		}
	}

	return -1
}

func BinarySearchRecursive(arr []int, low, high, target int) int {
	if high >= low {
		mid := ((high - low) / 2) + low
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			return BinarySearchRecursive(arr, low, mid-1, target)
		}

		return BinarySearchRecursive(arr, mid+1, high, target)
	}

	return -1
}
