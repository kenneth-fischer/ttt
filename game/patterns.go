package game

import (
	"strings"
)

func replaceEach(text, target string, patterns ... string) []string {
	results := []string{}

	for _, pattern := range patterns {
		result := text

		for i := 0; i < len(pattern); i++ {
			replacement := string(pattern[i])
			result = strings.Replace(result, target, replacement, 1) 
		}

		results = append(results, result)
	}
	return results
}

func Combine(s1 string, occurrences1 int, s2 string, occurrences2 int) []string {
	results := []string{}

	if occurrences1 < 0 || occurrences2 < 0 || occurrences1 + occurrences2 == 0 {
		return results
	}

	if occurrences1 == 0 {
		result := strings.Repeat(s2, occurrences2)
		results = append(results, result)
		return results
	}

	if occurrences2 == 0 {
		result := strings.Repeat(s1, occurrences1)
		results = append(results, result)
		return results
	}

        for _, partial := range Combine(s1, occurrences1-1, s2, occurrences2) {
		result := s1 + partial
		results = append(results, result)
	}

	for _, partial := range Combine(s1, occurrences1, s2, occurrences2-1) {
		result := s2 + partial
		results = append(results, result)
	}
	return results
}
