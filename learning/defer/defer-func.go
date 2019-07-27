package main

import "fmt"

func finished() {
	fmt.Println("Finished finding largest number.")
}
func largest(nums []int) {

	defer finished()

	fmt.Println("Started finding largest number")

	max := nums[0]

	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in ", nums, "is", max)

}
func deferFunc() {
	num := []int{78, 108, 2, 8, 1500, 547, 39}
	largest(num)

}
