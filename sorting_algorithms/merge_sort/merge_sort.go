package merge_sort

func SortArray(arr []int) []int {
	return MergeSort(arr)
}

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	arr1 := arr[:mid]
	arr2 := arr[mid:]

	arr1 = MergeSort(arr1)
	arr2 = MergeSort(arr2)

	return Merge(arr1, arr2)
}

// takes in 2 sorted arrays and merges them into one new sorted array
func Merge(arr1, arr2 []int) []int {
	newArr := make([]int, 0, len(arr1)+len(arr2))

	i := 0 // track arr1 index
	j := 0 // track arr2 index

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] > arr2[j] {
			newArr = append(newArr, arr2[j])
			j++
		} else {
			newArr = append(newArr, arr1[i])
			i++
		}
	}

	for l := i; l < len(arr1); l++ {
		newArr = append(newArr, arr1[l])
	}

	for l := j; l < len(arr2); l++ {
		newArr = append(newArr, arr2[l])
	}

	return newArr
}
