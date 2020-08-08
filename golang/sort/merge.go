package sort

// MergeSort using merge sort algorithm
func MergeSort(nums []int) []int {
	return splitMerge(nums)
}

func splitMerge(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := splitMerge(nums[:mid])
	right := splitMerge(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	if len(left) == 0 {
		return right
	}
	if len(right) == 0 {
		return left
	}

	var mergedSlice []int
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			mergedSlice = append(mergedSlice, left[l])
			l++
		} else {
			mergedSlice = append(mergedSlice, right[r])
			r++
		}
	}

	for l < len(left) {
		mergedSlice = append(mergedSlice, left[l])
		l++
	}
	for r < len(right) {
		mergedSlice = append(mergedSlice, right[r])
		r++
	}

	return mergedSlice
}
