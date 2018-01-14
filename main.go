package main

import (
	"brianj-za/cs-fundamentals/sort"
	"log"
	"math/rand"
)

func main() {
	log.Println("Starting...")

	r := rand.New(rand.NewSource(99))

	arr := r.Perm(20000)
	arr2 := sort.Bubble(arr)
	log.Println(arr2[:20])

	r = rand.New(rand.NewSource(99))

	arr = r.Perm(20000)
	arr2 = sort.Insertion(arr)
	log.Println(arr2[:20])

	arr = r.Perm(20000)
	arr = sort.Merge(arr)
	log.Println(arr[:20])
}
