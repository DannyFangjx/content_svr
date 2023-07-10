package utils

import (
	"sort"
)

func StrStrMapKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for _k := range m {
		keys = append(keys, _k)
	}
	sort.Strings(keys)
	return keys
}

func StrIMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for _k := range m {
		keys = append(keys, _k)
	}
	sort.Strings(keys)
	return keys
}
