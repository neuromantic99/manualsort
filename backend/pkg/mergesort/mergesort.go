package mergesort

import (
	"fmt"

	"bufio"

	"os"
)

func BeginSort() {

	// Slice literal to sort
	var toSort = []int{5, 4, 6, 9, 100}

	fmt.Println(toSort)
	mergeSort(toSort, len(toSort))
	fmt.Println(toSort)

}

func mergeSort(toSort []int, n int) {

	if n < 2 {
		return
	}

	// Find the middle point of the array
	middle := n / 2

	leftSlice := make([]int, middle)
	rightSlice := make([]int, n-middle)

	copy(leftSlice, toSort[0:middle])
	copy(rightSlice, toSort[middle:n])

	mergeSort(leftSlice, middle)
	mergeSort(rightSlice, n-middle)

	merge(toSort, leftSlice, rightSlice, middle, n-middle)

}

func merge(toSort []int, leftSlice []int, rightSlice []int, left int, right int) {

	var i, j, k int = 0, 0, 0

	for {

		if i >= left || j >= right {
			break
		}

		fmt.Printf("Is ")
		fmt.Println(leftSlice[i])
		fmt.Printf("bigger than ")
		fmt.Println(rightSlice[j])

		reader := bufio.NewReader(os.Stdin)
		char, _, _ := reader.ReadRune()

		if char != 'y' {

			toSort[k] = leftSlice[i]
			i++

		} else {

			toSort[k] = rightSlice[j]
			j++
		}

		k++
	}

	if i < left {
		copy(toSort[k:], leftSlice[i:])
	} else if j < right {
		copy(toSort[k:], rightSlice[j:])
	}

}

