package algorithms

import "fmt"


func BinarySearch () {
	
	haystack  := []int{-11,-3,3,4,5,6,7,8,9,10,12,17,43,67,86,96,100,101,102,103,104,105,105,106,120,140,504,505,1014,1200,4000,4001,39999}
	needle := 140
	idx := binarySearch(needle, haystack);
	if idx == -1 {
		fmt.Printf("Couldn't find %d in the haystack.\n", needle)
	} else {
		fmt.Printf("Found the needle at index %d.\n", idx)
	}

}

func binarySearch(needle int, haystack []int) int {

	var right, left int

	left  = 0
	right = len(haystack) - 1
	
	if haystack[right] < needle || haystack[left] > needle { return -1 }
	if haystack[right] == haystack[left] { 
		if haystack[right] == needle { return right }
		return -1
	}

	for right >= left {
		median := (left + right) / 2
		if haystack[median] == needle { return median }
		if haystack[median] >  needle { right = median - 1 } else { left = median + 1 }
	}

	if left == len(haystack) || haystack[left] != needle { return -1 }
	return right
}