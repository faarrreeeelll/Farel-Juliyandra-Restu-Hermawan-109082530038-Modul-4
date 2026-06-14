package main

import "fmt"

const NMax int = 1000000

type ArrayRumah [NMax]int

func SelectionSort(arr *ArrayRumah, n int) {
	var i, j, minIdx, temp int

	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		temp = arr[minIdx]
		arr[minIdx] = arr[i]
		arr[i] = temp
	}
}

func main() {
	var n, m, i, j int
	var rumah ArrayRumah
	var pertama bool

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&m)

		for j = 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		SelectionSort(&rumah, m)

		pertama = true

		for j = 0; j < m; j++ {
			if rumah[j]%2 != 0 {
				if !pertama {
					fmt.Print(" ")
				}
				fmt.Print(rumah[j])
				pertama = false
			}
		}

		for j = m - 1; j >= 0; j-- {
			if rumah[j]%2 == 0 {
				if !pertama {
					fmt.Print(" ")
				}
				fmt.Print(rumah[j])
				pertama = false
			}
		}
		fmt.Println()
	}
}