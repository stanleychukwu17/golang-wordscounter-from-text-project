package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

func main() {
	start := time.Now() // Record the start time

	// next: open the file
	file, err := os.Open("great-gatsby.txt")
	if err != nil {
		fmt.Printf("Error: %v \n", err)
		os.Exit(1)
	}
	defer file.Close() // close the file when we are done

	// next: read each words in the file loaded and tell us how many times a word appear in the document
	wordsCount, err := read_the_frequency_of_words(file)
	if err != nil {
		fmt.Printf("Error: %v \n", err)
		os.Exit(1)
	}

	// next: sort the words from the highest to the lowest
	sort_the_words(wordsCount)

	elapsed := time.Since(start) // Calculate elapsed time
	fmt.Printf("Time taken: %s\n", elapsed)
}

// next: read the words in the file and tell us how many times a word appear in the document
func read_the_frequency_of_words(f *os.File) (map[string]int, error) {
	wordsCount := make(map[string]int) // map to store the words and their frequency

	scanner := bufio.NewScanner(f) // create a new scanner using the

	// loop through each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		each_word_into_slice := strings.Split(line, " ") // split the line into a []string (i.e a slice containing each of the words per line)

		// loop through each word
		for _, eachWord := range each_word_into_slice {
			word := strings.ToLower(eachWord)
			word = strings.TrimSpace(word) // remove leading and trailing spaces

			re := regexp.MustCompile("[^a-z0-9']")       // Matches characters not in the allowed set
			cleanedWord := re.ReplaceAllString(word, "") // Replace all matches with empty string

			// if the word is not empty the increase the word count in the map
			if lenWord := len(cleanedWord); lenWord > 0 {
				wordsCount[cleanedWord]++
			}
		}
	}

	// check for errors from the scanner
	if err := scanner.Err(); err != nil {
		return wordsCount, err
	}

	return wordsCount, nil
}

// next: sort the words in the map from the highest to the lowest and prints it out
func sort_the_words(wordsCount map[string]int) {
	groupFrequency := make([]int, 0)         // slice to store the each numbers
	frequencyStore := make(map[int][]string) // map to store each of the words that appear for a frequency

	// loops through the each of the words and their count
	for _, word_appearX := range wordsCount {
		// if the frequency is not in the map then add it
		if _, ok := frequencyStore[word_appearX]; !ok {
			frequencyStore[word_appearX] = []string{}             // add the frequency to the map and initializes a slice to store each of the words that has this frequency
			groupFrequency = append(groupFrequency, word_appearX) // add the frequency to the slice
		}
	}

	// sort the slice of frequencies so when can have the top numbers first
	sort.Slice(groupFrequency, func(i, j int) bool {
		return groupFrequency[i] > groupFrequency[j]
	})

	// loops through the each of the words and add them to the right frequency store
	for word, frequency := range wordsCount {
		frequencyStore[frequency] = append(frequencyStore[frequency], word)
	}

	// print out the words and their frequency starting from the words with the most frequency
	for _, frequency := range groupFrequency {
		// grab each word from their frequency store and print it out with the frequency
		for _, word := range frequencyStore[frequency] {
			fmt.Printf("%s: %d \n", word, frequency)
		}
	}
}
