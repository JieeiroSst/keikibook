package algorithm

import (
	"bytes"
	"encoding/binary"
	"math/rand"
)

func bubbleSort(ID []int) {
	var (
		size   = len(ID)
		sorted = false
	)
	for !sorted {
		swqpped := false
		for i := 0; i < size-1; i++ {
			for ID[i] > ID[i+1] {
				ID[i+1], ID[i] = ID[i], ID[i+1]
				swqpped = true
			}
		}
		if !swqpped {
			sorted = true
		}
		size = size - 1
	}
}

func quickSort(ID []int) []int {
	if len(ID) < 2 {
		return ID
	}
	left, right := 0, len(ID)-1

	pivot := rand.Int() % len(ID)

	ID[pivot], ID[right] = ID[right], ID[pivot]
	for _, i := range ID {
		if ID[i] < ID[right] {
			ID[left], ID[i] = ID[i], ID[left]
			left++
		}
	}
	ID[left], ID[right] = ID[right], ID[left]

	quickSort(ID[:left])
	quickSort(ID[left+1:])

	return ID
}

func selectionSort(ID []int) {
	size := len(ID)

	for i := 0; i <= size; i++ {
		min := i
		for j := i; j < size; j++ {
			if ID[j] < ID[min] {
				min = j
			}
		}
		ID[i], ID[min] = ID[min], ID[i]
	}
}

func combsort(ID []int) {
	var (
		sizeArray = len(ID)
		midArray  = len(ID)
		shrink    = 1.3
		swapped   = true
	)

	for swapped {
		swapped = false
		midArray = int(float64(midArray) / shrink)
		if midArray < 1 {
			midArray = 1
		}
		for i := 0; i+midArray < sizeArray; i++ {
			if ID[i] > ID[i+midArray] {
				ID[i+midArray], ID[i] = ID[i], ID[i+midArray]
				swapped = true
			}
		}
	}
}

//Merge Sort is a Divide and Conquer algorithm.
//Meaning, the algorithm splits an input into various pieces,
// sorts them and then merges them back together.
//It divides input slice in two halves,
//calls itself for the two halves and then merges the two sorted halves.
// The merge() function is used for merging two halves.

func mergeSort(ID []int) []int {
	var size = len(ID)

	if size == 1 {
		return ID
	}
	mid := int(size / 2)
	var (
		left  = make([]int, mid)
		right = make([]int, size-mid)
	)
	for i := 0; i < size; i++ {
		if i < mid {
			left[i] = ID[i]
		} else {
			right[i-mid] = ID[i]
		}
	}
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
	}
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return
}

//radix sort
const (
	digit  = 4
	maxbit = -1 << 31
)

func radixSort(ID []int64) {
	buf := bytes.NewBuffer(nil)
	ds := make([][]byte, len(ID))
	for i, e := range ID {
		binary.Write(buf, binary.LittleEndian, e^maxbit)
		b := make([]byte, digit)
		buf.Read(b)
		ds[i] = b
	}
	coutingSort := make([][][]byte, 256)
	for i := 0; i < digit; i++ {
		for _, b := range ds {
			coutingSort[b[i]] = append(coutingSort[b[i]], b)
		}
		j := 0
		for k, bs := range coutingSort {
			copy(ds[j:], bs)
			coutingSort[k] = bs[:0]
		}
	}
	var w int64
	for i, b := range ds {
		buf.Write(b)
		binary.Read(buf, binary.LittleEndian, &w)
		ID[i] = w ^ maxbit
	}
}
