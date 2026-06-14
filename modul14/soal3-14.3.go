package main

import "fmt"

const NMax int = 1000000

type ArrayData [NMax]int

func InsertionSort(arr *ArrayData, n int) {
	var i, j, key int
	for i = 1; i < n; i++ {
		key = arr[i]
		j = i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var data ArrayData
	var n, val int

	n = 0
	for {
		fmt.Scan(&val)

		if val == -5313 || val < 0 {
			break
		}

		if val == 0 {
			if n > 0 {
				InsertionSort(&data, n)
				if n%2 != 0 {
					fmt.Println(data[n/2])
				} else {
					fmt.Println((data[n/2-1] + data[n/2]) / 2)
				}
			}
		} else {
			data[n] = val
			n++
		}
	}
}