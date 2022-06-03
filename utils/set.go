package utils

import (
	"sort"
	"strings"
)

// RemoveDuplicate 去除string数组中重复的元素
func RemoveDuplicate(list ...string) []string {
	sort.Strings(list)
	i := 0
	for j := 1; j < len(list); j++ {
		if strings.Compare(list[i], list[j]) == -1 {
			i++
			list[i] = list[j]
		}
	}
	return list[:i+1]
}
