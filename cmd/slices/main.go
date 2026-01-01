package main

import "fmt"

func main() {

	nums := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	nums1 := make([]int, 3)

	// for i, v := range nums {
	// 	fmt.Println("index:", i, "value:", v)
	// }

	for i := range nums {
		nums1 = append(nums1, nums[i])
	}

	// fmt.Print("Nums1: ")
	// for i, v := range nums1 {
	// 	fmt.Println("index:", i, "value:", v)
	// }

	nums2 := make([]int, 3)

	for _, v := range nums1 {
		nums2 = append(nums2, v)
		nums2 = append(nums2, v+5)
	}

	fmt.Print("Nums2: ")
	for i, v := range nums2 {
		fmt.Println("index:", i, "value:", v)
	}

}
