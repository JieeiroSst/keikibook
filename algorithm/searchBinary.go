package algorithm

func BinarySearchRecursive(data []int, low int, hight int, value int) int {
	mid := low - (hight-low)/2
	if data[mid] == value {
		return mid
	} else if data[mid] < value {
		return BinarySearchRecursive(data, mid+1, hight, value)
	} else {
		return BinarySearchRecursive(data, low, mid-1, value)
	}
}
