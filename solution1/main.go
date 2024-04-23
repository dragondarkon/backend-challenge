package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func maxPathSum() {

	// Uncomment to test
	// arr := [][]int{
	// 	{59},
	// 	{73, 41},
	// 	{52, 40, 53},
	// 	{26, 53, 6, 34},
	// }

	arr := [][]int{}
	data, _ := os.ReadFile("./files/hard.json")
	json.Unmarshal(data, &arr)

	if len(arr) == 0 {
		fmt.Println(0)
		return
	}

	for i := len(arr) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			var left = arr[i+1][j]
			fmt.Print(left)
			fmt.Print(",")
			var right = arr[i+1][j+1]
			fmt.Print(right)
			fmt.Println()
			if left >= right {
				arr[i][j] += left
			} else {
				arr[i][j] += right
			}
			fmt.Println()
		}
	}
	fmt.Println(arr[0][0])
}

func main() {
	maxPathSum()
}
