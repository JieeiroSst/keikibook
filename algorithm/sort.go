package algorithm

import (
	"fmt"
	"sort"
)

func handleSort(x int) {
	ID := []int{}
	i := sort.Search(len(ID), func(i int) bool { return ID[i] >= x })

	if i < len(ID) && ID[i] == x {
		fmt.Println("found")
	} else {
		fmt.Println("found'n")
	}
}
