package main

func removeElement_1(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)

}

func removeElement_2(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			lookup := false
			for j := i; j < len(nums); j++ {
				if nums[j] != val {
					nums[i], nums[j] = nums[j], nums[i]
					lookup = true
					break
				}
			}
			if !lookup {
				return i
			}
		}
	}
	return len(nums)
}

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
