package quick_sort

func SortArray(arr []int) []int {
	QuickSort(arr, 0, len(arr)-1)
	return arr
}

func QuickSort(arr []int, low, high int) {
	if low >= high {
		return
	}

	pivotIndex := partition(arr, low, high)
	QuickSort(arr, low, pivotIndex-1)
	QuickSort(arr, pivotIndex+1, high)
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if pivot > arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return i
}
