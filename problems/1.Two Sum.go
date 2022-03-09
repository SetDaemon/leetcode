package problems

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		expect := target - num
		if value, ok := m[expect]; ok {
			return []int{i, value}
		}
		m[num] = i
	}
	return nil
}
