package sort

import (
	"log"
	"time"
)

func Merge(arr []int) []int {
	defer timeTrack(time.Now(), "Merge Sort")
	return splitAndMergeM(arr)
}

func splitAndMergeM(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	middle := len(arr) / 2
	return merge(splitAndMergeM(arr[:middle]), splitAndMergeM(arr[middle:]))
}

func merge(arrLeft []int, arrRight []int) []int {
	maxLen := len(arrLeft)
	if len(arrRight) > maxLen {
		maxLen = len(arrRight)
	}
	results := make([]int, 0, maxLen*2)
	for len(arrRight) > 0 && len(arrLeft) > 0 {
		if arrRight[0] <= arrLeft[0] {
			results = append(results, arrRight[0])
			arrRight = arrRight[1:]
		} else {
			results = append(results, arrLeft[0])
			arrLeft = arrLeft[1:]
		}
	}

	results = append(results, arrLeft...)
	results = append(results, arrRight...)

	return results
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
