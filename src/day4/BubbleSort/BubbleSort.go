package bubblesort

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			curr, next := arr[j], arr[j+1]
			if curr > next {
				arr[j], arr[j+1] = next, curr
			}
		}
	}
}
