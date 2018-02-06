package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 && n != 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	n1 := make([]int, len(nums1))
	copy(n1, nums1)

	var i, j int
	for x := 0; x < len(nums1); x++ {
		if i >= m {
			nums1[x] = nums2[j]
			j++
			continue
		}
		if j >= n {
			nums1[x] = n1[i]
			i++
			continue
		}
		if nums2[j] > n1[i] {
			nums1[x] = n1[i]
			i++
		} else {
			nums1[x] = nums2[j]
			j++
		}
	}
}
