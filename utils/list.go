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

// RemoveOne 去除string数组的目标字符串,没有则返回原数组的深拷贝
func RemoveOne(src []string, target string) []string {
	list := make([]string, len(src))
	copy(list, src)
	sort.Strings(list)
	i := sort.SearchStrings(list, target)
	if i >= len(list) || list[i] != target {
		return list
	}
	copy(list[i:], list[i+1:])
	list = list[:len(list)-1]
	return list
}
