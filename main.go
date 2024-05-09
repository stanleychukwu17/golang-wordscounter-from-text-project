package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("great-gatsby.txt")
	if err != nil {
		fmt.Printf("Error: %v \n", err)
		os.Exit(1)
	}
	defer file.Close()

	// next: read each words in the file loaded
	wordsCount, err := read_the_frequency_of_words(file)
	if err != nil {
		fmt.Printf("Error: %v \n", err)
		os.Exit(1)
	}

	// fmt.Print(wordsCount)
	// next: sort the words in the map
	sort_the_words(wordsCount)
}

func read_the_frequency_of_words(f *os.File) (map[string]int, error) {
	wordsCount := make(map[string]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		eachWord := strings.Split(line, " ")

		for _, eachText := range eachWord {
			word := strings.ToLower(eachText)
			word = strings.TrimSpace(word)

			re := regexp.MustCompile("[^a-z0-9']") // Matches characters not in the allowed set
			cleanedWord := re.ReplaceAllString(word, "")

			if lenWord := len(cleanedWord); lenWord > 0 {
				wordsCount[cleanedWord]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return wordsCount, err
	}

	return wordsCount, nil
}

// next: sort the words in the map
func sort_the_words(wordsCount map[string]int) []string {
	uniqueNumbersSlice := make([]int, 0)
	uniqueNumbersMap := make(map[int][]string)

	for _, word_appearX := range wordsCount {
		if _, ok := uniqueNumbersMap[word_appearX]; !ok {
			uniqueNumbersMap[word_appearX] = []string{}
			uniqueNumbersSlice = append(uniqueNumbersSlice, word_appearX)
		}
	}

	sort.Slice(uniqueNumbersSlice, func(i, j int) bool {
		return uniqueNumbersSlice[i] > uniqueNumbersSlice[j]
	})

	for word, word_appearX := range wordsCount {
		uniqueNumbersMap[word_appearX] = append(uniqueNumbersMap[word_appearX], word)
	}

	// fmt.Println(uniqueNumbersMap)
	for _, word_appearX := range uniqueNumbersSlice {
		for _, word := range uniqueNumbersMap[word_appearX] {
			fmt.Printf("%s: %d \n", word, word_appearX)
		}
	}

	return []string{}
}
