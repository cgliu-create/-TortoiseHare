package algo

// FindDuplicate (floyds algorithem)
// moves two pointers at different speeds through the sequence of values
// until they both point to equal values
// given array with integers >= 1 and <= size of the array
func FindDuplicate(nums []int) int {
	t := nums[0]
	h := nums[0]
	for {
		t = nums[t]
		h = nums[nums[h]]
		if t == h {
			break
		}
	}
	p1 := nums[0]
	p2 := t
	for p1 != p2 {
		p1 = nums[p1]
		p2 = nums[p2]
	}
	return nums[p1]
}
