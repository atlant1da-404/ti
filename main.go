package main

import "fmt"

func main() {

	//arr := []int{363374326, 364147530, 61825163, 1073065718, 1281246024, 1399469912, 428047635, 491595254, 879792181, 1069262793}

	//arr := []int{363374326, 364147530, 61825163, 1073065718, 1281246024, 1399469912, 428047635, 491595254, 879792181, 1069262793}
	//arr := []int{1, 5, 3, 4, 2}

	arr := []int32{1, 2, 3, 4}
	var k int32 = 1

	fmt.Println(pairs(k, arr))
}

func pairs(k int32, arr []int32) int32 {

	var count int32

	if len(arr) == 0 {
		return 0
	}

	for i := range arr {

		j := i + 1

		for j < len(arr) {

			if arr[j]-arr[i] == k || arr[i]-arr[j] == k {
				count++
			}

			j++
		}
	}
	return count
}
