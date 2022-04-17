package utils

import "math"

// 简单的切片分页功能
func SlicePage(page, pageSize, total int) (sliceStart, sliceEnd int) {
	if page < 0 {
		page = 1
	}

	if pageSize < 0 {
		pageSize = 10
	}

	if pageSize > total {
		return 0, total
	}

	pageCount := int(math.Ceil(float64(total) / float64(pageSize)))
	if page > pageCount {
		return 0, 0
	}

	sliceStart = (page - 1) * pageSize
	sliceEnd = sliceStart + pageSize

	if sliceEnd > total {
		sliceEnd = total
	}

	return sliceStart, sliceEnd
}
