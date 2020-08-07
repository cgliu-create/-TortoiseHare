package algo

// FindDuplicate (floyds algorithem)
// moves two pointers at different speeds through the sequence of values
// until they both point to equal values
// given array of size n with values that conveniently range from 0 to n-1. *cough cough index*
func FindDuplicate(nums []int) int {
	if len(nums) <= 1 {
		return -1
	}
	t := nums[0]
	h := nums[nums[0]]
	for t != h {
		t = nums[t]
		h = nums[nums[h]]
	}
	p := 0
	for p != t {
		p = nums[p]
		t = nums[t]
	}
	return p
}
