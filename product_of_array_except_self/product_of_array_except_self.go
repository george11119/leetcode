package product_of_array_except_self

// O(n) time complexity, O(1) extra space other than return array
// damn prefix sum is magic
func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	product := 1
	for i := 0; i < len(nums); i++ {
		res[i] = product
		product *= nums[i]
	}

	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= product
		product *= nums[i]
	}

	return res
}

// brute force
//func ProductExceptSelf(nums []int) []int {
//	res := make([]int, len(nums))
//
//	for i := 0; i < len(nums); i++ {
//		sum := 1
//		for j := 0; j < len(nums); j++ {
//			if i != j {
//				sum *= nums[j]
//			}
//		}
//		res[i] = sum
//	}
//
//	return res
//}

// O(n) time, O(n) space
//func ProductExceptSelf(nums []int) []int {
//	prefix := make([]int, len(nums)+2)
//	postfix := make([]int, len(nums)+2)
//
//	sum := 1
//	prefix[0] = sum
//	prefix[len(prefix)-1] = sum
//	for i := 0; i < len(nums); i++ {
//		sum *= nums[i]
//		prefix[i+1] = sum
//	}
//
//	sum = 1
//	postfix[0] = sum
//	postfix[len(postfix)-1] = sum
//	for i := len(nums) - 1; i >= 0; i-- {
//		sum *= nums[i]
//		postfix[i+1] = sum
//	}
//
//	for i := 1; i < len(nums)+1; i++ {
//		product := prefix[i-1] * postfix[i+1]
//		nums[i-1] = product
//	}
//
//	return nums
//}
