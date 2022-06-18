
## Note

##### Sort
```Golang
	nums := 2934
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	// nums = 9432
```