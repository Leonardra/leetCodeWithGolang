package easy

func twoSums (nums []int, target int) []int {
	var result = []int{0,0}
	mapper := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		remainder :=  target - nums[i]
		_, ok := mapper[remainder]
		if  ok {
			result[0] = mapper[remainder]
			result[1] = i
		}else{
			mapper[nums[i]] = i
		}
	}
	return result
}
