package algorithm

import (
	"fmt"
	"sort"
)

func handleSort() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := 3
	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })

	if i < len(a) && a[i] == x {
		fmt.Println("done")
	} else {
		fmt.Println("no done")
	}
}
