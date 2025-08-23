package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

func Top10(text string) []string {
	words := strings.Fields(text)

	counts := map[string]int{}
	for _, w := range words {
		counts[w]++
	}

	allWords := make([]string, 0, len(counts))
	for w := range counts {
		allWords = append(allWords, w)
	}

	sort.Slice(allWords, func(i, j int) bool {
		if counts[allWords[i]] == counts[allWords[j]] {
			return allWords[i] < allWords[j]
		}
		return counts[allWords[i]] > counts[allWords[j]]
	})

	if len(allWords) > 10 {
		return allWords[:10]
	}
	fmt.Println(allWords)
	return allWords
}
