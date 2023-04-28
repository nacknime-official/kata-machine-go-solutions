package quicksort

func qs(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)

	qs(arr, lo, pivotIdx)
	qs(arr, pivotIdx+1, hi)
}

func partition(arr []int, lo, hi int) int {
	pivotIdx := (lo + hi) / 2
	pivot := arr[pivotIdx]
	i, j := lo, hi

	for {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}

		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func QuickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}
