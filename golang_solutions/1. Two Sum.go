package golang_solutions

func twoSum(nums []int, target int) []int {
	var holder []int
	var ans []int

	for i, v := range nums {
		for j, k := range holder {
			if target-v == k {
				ans = append(ans, i, j)
				break
			}
		}
		holder = append(holder, v)
	}
	return ans
}
