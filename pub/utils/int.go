package utils

import (
	"fmt"
	"strconv"
)

// 支持int和string 转化为int64
func Interface2Int64(src interface{}) (int64, error) {
	srcStr := InterfaceToString(src)
	valInt, err := strconv.ParseInt(srcStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return valInt, err
}

func IsInt64Repeat(items []int64) bool {
	m := make(map[int64]struct{}, len(items))
	for _, item := range items {
		if _, ok := m[item]; ok {
			return true
		}
		m[item] = struct{}{}
	}
	return false
}

func Int642Uint32(items []int64) []uint32 {
	result := make([]uint32, 0, len(items))
	for _, i := range items {
		result = append(result, uint32(i))
	}
	return result
}

func Int(i int) *int {
	return &i
}

func Int64(i int64) *int64 {
	return &i
}

func Int64In(list []int64, target int64) bool {
	for _, i := range list {
		if target == i {
			return true
		}
	}
	return false
}

func Int64SliceDiffer(source, target []int64, deduplicate bool) []int64 {
	if len(target) == 0 {
		return source
	}
	result := make([]int64, 0, len(source))
	if len(source) == 0 {
		return result
	}
	targetMap := make(map[int64]bool, len(target))
	for _, v := range target {
		targetMap[v] = true
	}
	for _, v := range source {
		_, ok := targetMap[v]
		if !ok && deduplicate && !Int64In(result, v) {
			result = append(result, v)
		} else if !ok && !deduplicate {
			result = append(result, v)
		}
	}
	return result
}

func Int64Set(nums []int64) []int64 {
	intMap := make(map[int64]bool)
	for _, _n := range nums {
		intMap[_n] = true
	}
	result := make([]int64, 0, len(intMap))
	for k := range intMap {
		result = append(result, k)
	}
	return result
}

func Int642StrList(nums []int64) string {
	list := JoinInt64s(nums, ", ")
	return fmt.Sprintf("%s%s%s", "[", list, "]")
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IInt64(value interface{}) int64 {
	switch t := value.(type) {
	case int:
		return int64(t)
	case int8:
		return int64(t)
	case int16:
		return int64(t)
	case int32:
		return int64(t)
	case uint:
		return int64(t)
	case uint8:
		return int64(t)
	case uint16:
		return int64(t)
	case uint32:
		return int64(t)
	case uint64:
		return int64(t)
	}
	return value.(int64)
}

func IInt(value interface{}) int {
	switch t := value.(type) {
	case int8:
		return int(t)
	case int16:
		return int(t)
	case int32:
		return int(t)
	case int64:
		return int(t)
	case uint:
		return int(t)
	case uint8:
		return int(t)
	case uint16:
		return int(t)
	case uint32:
		return int(t)
	case uint64:
		return int(t)
	}
	return value.(int)
}
