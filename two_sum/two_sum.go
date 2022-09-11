package two_sum

func TwoSumBrute(nums []int, target int) []int {
	for i, v := range nums {
		if i == (len(nums) - 1) {
			break
		}

		offset := i + 1
		for j, u := range nums[offset:] {
			if v+u == target {
				return []int{i, j + offset}
			}
		}
	}

	return []int{}
}

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}

	for i, v := range nums {
		if j, ok := m[target-v]; ok && j != i {
			return []int{i, j}
		}
	}

	return []int{}
}
