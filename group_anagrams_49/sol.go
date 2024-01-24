package group_anagrams_49

import "sort"

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		s := []rune(strs[i])
		sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		strMap[sortedStr] = append(strMap[sortedStr], strs[i])
	}
	result := make([][]string, len(strMap))
	idx := 0
	for _, v := range strMap {
		result[idx] = v
		idx++
	}
	return result
}
