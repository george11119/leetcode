package two_sum

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int) // should be value: index
	for i, v := range nums {
		difference := target - v
		j, ok := myMap[difference]
		if ok {
			return []int{j, i}
		}
		myMap[v] = i
	}
	return []int{}
}
