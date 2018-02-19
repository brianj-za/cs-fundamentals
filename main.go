package main

import (
	"log"
	"math/rand"

	"bitbucket.org/brianj-za/cs-fundamentals/sort"
)

func main() {
	log.Println("Starting...")

	r := rand.New(rand.NewSource(99))
	arr := r.Perm(20000)
	arr2 := sort.Bubble(arr)

	// r = rand.New(rand.NewSource(99))
	// arr = r.Perm(20000)
	// arr3 := sort.Insertion(arr)
	// if len(arr) != len(arr3) {
	// 	panic("I broke!")
	// }

	r = rand.New(rand.NewSource(99))
	arr = r.Perm(20000)
	arr4 := sort.Merge(arr)
	if len(arr) != len(arr4) {
		panic("I broke!")
	}

	// r = rand.New(rand.NewSource(99))
	// arr = r.Perm(20000)
	// arr5 := sort.Quick(arr)
	// if len(arr) != len(arr5) {
	// 	panic("I broke!")
	// }

	// As a test, take sorted case
	log.Println("\nSorted cases")
	arr2 = sort.Bubble(arr2)
	// arr3 = sort.Insertion(arr3)
	arr4 = sort.Merge(arr4)
	// arr5 = sort.Quick(arr5)

	log.Println("\nReverse Sorted cases")
	arr2 = sort.Bubble(reverse(arr2))
	// arr3 = sort.Insertion(reverse(arr3))
	arr4 = sort.Merge(reverse(arr4))
	// arr5 = sort.Quick(reverse(arr5))
}

func reverse(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
