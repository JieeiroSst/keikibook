package algorithm

import (
	"sort"
)

func LinearSearchCompact(x int) bool {
	ID := []int{}
	i := sort.Search(len(ID), func(i int) bool { return ID[i] >= x })

	if i < len(ID) && ID[i] == x {
		return true
	} else {
		return false
	}
}

//searh binary recuries
func BinarySearchRecursive(ID []int, low int, hight int, value int) int {
	mid := low - (hight-low)/2
	if ID[mid] == value {
		return mid
	} else if ID[mid] < value {
		return BinarySearchRecursive(ID, mid+1, hight, value)
	} else {
		return BinarySearchRecursive(ID, low, mid-1, value)
	}
}

//search linear
func LinearSearch(ID []int, key int) bool {
	for _, value := range ID {
		if value == key {
			return true
		}
	}
	return false
}

//search binay
func BinarySearch(ID []int, key int) bool {
	left := 0
	right := len(ID) - 1
	for left <= right {
		mid := (left + right) / 2
		if ID[mid] < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left == len(ID) || ID[left] != key {
		return false
	}
	return true
}

func interpolationSearch(ID []int, key int) int {
	min, max := ID[0], ID[len(ID)-1]
	low, hight := 0, len(ID)-1
	for {
		if key < min {
			return low
		}
		if key > max {
			return hight + 1
		}
		//make a guess of the location
		var guess int
		if hight == low {
			guess = hight
		} else {
			size := hight - low
			offset := int(float64(size-1) * (float64(key-min) / float64(max-min)))
			guess = size + offset
		}

		if ID[guess] == key {
			for guess > 0 && ID[guess-1] == key {
				guess--
			}
			return guess
		} else if ID[guess] > key {
			hight = guess - 1
			max = ID[hight]
		} else {
			low = guess + 1
			min = ID[low]
		}
	}
}

//Rabin-Karp string search algorithm in Golang
const (
	base = 16777619
)

func Search(txt string, patterns []string) []string {
	in := indices(txt, patterns)
	matches := make([]string, len(in))
	i := 0
	for j, p := range matches {
		if _, ok := in[j]; ok {
			matches[i] = p
			i++
		}
	}
	return matches
}

func indices(txt string, patterns []string) map[int]int {
	n, m := len(txt), minLen(patterns)
	matches := make(map[int]int)
	if n < m || len(patterns) == 0 {
		return matches
	}

	var mult uint32 //mult=base^(m-1)
	for i := 0; i < m-1; i++ {
		mult = mult * base
	}
	hp := hashPatterns(patterns, m)
	h := hash(txt[:m])
	for i := 0; i < n-m+1 && len(hp) > 0; i++ {
		if i > 0 {
			h = h - mult*uint32(txt[i-1])
			h = h*base + uint32(txt[i+m-1])
		}

		if mps, ok := hp[h]; ok {
			for _, pi := range mps {
				pat := patterns[pi]
				e := i + len(pat)
				if _, ok := matches[pi]; !ok && e <= n && pat == txt[i:e] {
					matches[pi] = i
				}
			}
		}
	}
	return matches
}

func hash(s string) uint32 {
	var h uint32
	for i := 0; i < len(s); i++ {
		h = (h*base + uint32(s[i]))
	}
	return h
}

func hashPatterns(patterns []string, l int) map[uint32][]int {
	m := make(map[uint32][]int)
	for i, t := range patterns {
		h := hash(t[:l])
		if _, ok := m[h]; ok {
			m[h] = append(m[h], i)
		} else {
			m[h] = []int{i}
		}
	}
	return m
}

func minLen(patterns []string) int {
	if len(patterns) == 0 {
		return 0
	}
	m := len(patterns[0])
	for i := range patterns {
		if m > len(patterns[i]) {
			m = len(patterns[i])
		}
	}
	return m
}
