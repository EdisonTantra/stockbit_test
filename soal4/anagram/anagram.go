package anagram

import (
	"sort"
	"strings"
)

type anagram struct {
	Key    string
	Length int
	Value  []string
}

func isAnagram(str1 string, str2 string) bool {
	countMap := make(map[string]int)
	var flag bool = true

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)
	str1 = strings.Replace(str1, " ", "", -1)
	str2 = strings.Replace(str2, " ", "", -1)

	if len(str1) == len(str2) {
		for _, r := range str1 {
			j := countMap[string(r)]

			if j == 0 {
				countMap[string(r)] = 1
			} else {
				countMap[string(r)] = j + 1
			}
		}

		for _, r := range str2 {
			j := countMap[string(r)]

			if j == 0 {
				countMap[string(r)] = 1
			} else {
				countMap[string(r)] = j + 1
			}
		}

		for _, value := range countMap {
			if value%2 != 0 {
				flag = false
				break
			}
		}
	} else {
		flag = false
	}

	return flag
}

func formatResponse(lookup map[string][]string) (result [][]string) {
	var ss []anagram
	for k, v := range lookup {
		ss = append(ss, anagram{k, len(v), v})
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Length == ss[j].Length {
			return ss[i].Key < ss[j].Key
		} else {
			return ss[i].Length > ss[j].Length
		}
	})

	idx := 0
	result = make([][]string, len(lookup))
	for _, kv := range ss {
		result[idx] = kv.Value
		idx++
	}

	return result
}

func Detect(arr []string) (result [][]string) {
	lookup := make(map[string][]string)

	for _, elem := range arr {
		newWord := true
		if len(lookup) > 0 {
			for key, _ := range lookup {
				if strings.ToLower(key) != strings.ToLower(elem) {
					if isAnagram(key, elem) {
						lookup[key] = append(lookup[key], elem)
						newWord = false
						break
					}
				} else {
					newWord = false
					break
				}
			}
		}
		if newWord {
			lookup[elem] = []string{elem}
		}
	}

	result = formatResponse(lookup)
	return result
}
