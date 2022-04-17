package utils

import (
	"testing"
)

func TestPagination(t *testing.T) {
	var (
		page     = 1
		pageSize = 10
	)

	arr := []struct {
		name          string
		size          int
		last_modified string
	}{
		{name: "l.png", size: 245701, last_modified: "2022-03-17 02:57:43"},
		{name: "t.png", size: 245702, last_modified: "2022-03-18 02:57:43"},
		{name: "z.png", size: 245703, last_modified: "2022-03-19 02:57:43"},
	}

	star, end := SlicePage(page, pageSize, len(arr))
	list := arr[star:end]

	if star != 1 && end != 3 {
		t.Errorf("star or end  Error")
	}

	if len(list) != 3 {
		t.Errorf("Pagination Error")
	}
}
