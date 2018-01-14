package sort

import "time"

func Insertion(arr []int) []int {
	defer timeTrack(time.Now(), "Insertion Sort")

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				t := arr[j]
				arr[j] = arr[i]
				arr[i] = t
			}
		}
	}

	return arr
}
