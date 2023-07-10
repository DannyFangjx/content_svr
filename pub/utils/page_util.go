package utils

import "math"

type PageResult struct {
	PageNum       uint64 `json:"page_num,omitempty"`
	PageSize      uint64 `json:"page_size,omitempty"`
	Total         uint64 `json:"total"`
	TotalPageSize uint64 `json:"total_page_size"`
}
type PageResultWithData struct {
	PageNum       uint64      `json:"page_num,omitempty"`
	PageSize      uint64      `json:"page_size,omitempty"`
	List          interface{} `json:"list"`
	Total         uint64      `json:"total"`
	TotalPageSize uint64      `json:"total_page_size"`
}

func CalculatePage(total uint64, pageSize uint64) uint64 {
	return uint64(math.Ceil(float64(total) / float64(pageSize)))
}
func CalculateOffSet(pageNum uint64, pageSize uint64) int {
	if pageNum == 0 {
		pageNum = 1
	}
	return int((pageNum - 1) * pageSize)
}
