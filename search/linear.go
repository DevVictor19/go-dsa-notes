package search

/*

# Linear Search Algorithm - O(n)

In Linear Search, we iterate over all the elements of the array and check if it the
current element is equal to the target element.

If we find any element to be equal to the target element, then return the index of
the current element. Otherwise, if no element is equal to the target element, then
return -1 as the element is not found.

## Time and complexity
- best case: O(1)
- worst case: O(n)
- average case: O(n)

## Auxiliary space
- O(1) as except for the variable to iterate through the list, no other variable is used.

## Applications of Linear Search Algorithm:
- Unsorted Lists
- Small Data Sets
- Searching Linked Lists
- Simple Implementation

## Disadvantages of Linear Search Algorithm:
- Linear search has a time complexity of O(N), which in turn makes it slow for large datasets.
- Not suitable for large arrays.

*/

func LinearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}
