package main

import "fmt"

func getSum(nums []int) int{
	var sum int = 0
	fmt.Printf()
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(getSum(nums))
}
