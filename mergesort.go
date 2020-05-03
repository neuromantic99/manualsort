package main

import "fmt"

func main() {

	// Slice literal to sort
	var toSort = []int{54, 56, 0, 99, 1, -1000, 200220, 89}

	fmt.Printf("Hi\n")

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

		if leftSlice[i] <= rightSlice[j] {

			toSort[k] = leftSlice[i]
			i++

		} else {

			toSort[k] = rightSlice[j]
			j++
		}
		k++
	}

	for i < left {

		toSort[k] = leftSlice[i]
		i++
		k++
	}

	for j < right {

		toSort[k] = rightSlice[j]
		j++
		k++
	}

}

