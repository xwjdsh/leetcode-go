package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	for ; i < len(nums1); i++ {
		for ; j < len(nums2); j++ {
			if nums2[j] < nums1[i] {
				nums1 = append(nums1[:i], append([]int{nums2[j]}, nums1[i:]...)...)
				i++
			} else {
				break
			}
		}
	}
	nums1 = append(nums1, nums2[j:]...)
	// log.Println(nums1)
	var result float64
	if len(nums1) != 0 {
		m := len(nums1) / 2
		if len(nums1)%2 == 0 {
			result = float64((nums1[m] + nums1[m-1])) / 2
		} else {
			result = float64(nums1[m])
		}
	}
	return result
}
