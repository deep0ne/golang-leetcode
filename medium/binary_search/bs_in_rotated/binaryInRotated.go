package main

func search(nums []int, target int) int {
	if len(nums) == 1 && nums[0] == target {
		return 0
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[l] == target {
			return l
		}

		if nums[r] == target {
			return r
		}

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > nums[r] {
			if nums[mid] > target && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
