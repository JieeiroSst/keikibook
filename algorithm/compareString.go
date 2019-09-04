package algorithm

import "fmt"

func max(more ...int) int {
	max := more[0]
	for _, elem := range more {
		if max < more[elem] {
			max = more[elem]
		}
	}
	return max
}

func longest(string1, string2 string) int {
	tab := make([][]int, len(string1)+1)
	for i := range tab {
		tab[i] = make([]int, len(string2)+1)
	}

	for i := 0; i <= len(string1); i++ {
		for j := 0; j <= len(string2); j++ {
			if i == 0 || j == 0 {
				tab[i][j] = 0
			} else if string1[i-1] == string2[j-1] {
				tab[i][j] = tab[i-1][j-1] + 1
				if i < len(string1) {
					fmt.Printf("%c", string1[i])

					i++
					j++
				}
			} else {
				tab[i][j] = max(tab[i-1][j], tab[i][j-1])
			}
		}
	}
	fmt.Println()
	return tab[len(string1)][len(string2)]
}
