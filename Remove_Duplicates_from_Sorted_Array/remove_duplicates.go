package rm_duplicates

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ptr := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[ptr] = nums[i]
			ptr++
		}
	}
	return ptr
}
