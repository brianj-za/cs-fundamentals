package sort

import "time"

func Quick(arr []int) []int {
	defer timeTrack(time.Now(), "Quick Sort")

	return splitAndMergeQ(arr)
}

func splitAndMergeQ(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)-1]
	left := make([]int, 0, len(arr)-1)
	right := make([]int, 0, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	arr = append(splitAndMergeM(left), pivot)
	arr = append(arr, splitAndMergeM(right)...)
	return arr
}
