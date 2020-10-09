package array

func Array2() [2]int {
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	return arr
}

func Array3() ([3]int, [3]string) {

	var intArr = [3]int{1, 2, 3}

	strArr := [3]string{"a", "b", "c"}

	return intArr, strArr
}
