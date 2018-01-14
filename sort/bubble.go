package sort

import "time"

func Bubble(arr []int) []int {
	defer timeTrack(time.Now(), "Bubble Sort")

	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] >= arr[i+1] {
				t := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = t
				swapped = true
			}
		}
	}
	return arr
}
