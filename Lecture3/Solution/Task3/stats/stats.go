package stats




func Largest(nums []int) int {
	if len(nums) == 0 {
		panic("slice is empty")
	}
	largest := nums[0]
	for _, num := range nums {
		if num > largest {		
			largest = num
		}
	}
	return largest
}
func Smallest(nums []int) int {
	if len(nums) == 0 {
		panic("slice is empty")
	}
	smallest := nums[0]
	for _, num := range nums {
		if num < smallest {

			smallest = num
		}
	}
	return smallest
}
func Average(nums []int) float64 {
	if len(nums) == 0 {
		panic("slice is empty")
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	average := float64(sum) / float64(len(nums))
	return average
}
