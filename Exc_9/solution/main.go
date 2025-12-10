package main

import (
	"bufio"
	"exc9/mapred"
	"fmt"
	"log"
	"os"
	"sort"
)

// Main function
func main() {
	// todo read file
	// Read file
	file, err := os.Open("res/meditations.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var text []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// todo run your mapreduce algorithm
	var mr mapred.MapReduce
	results := mr.Run(text)

	// todo print your result to stdout
	type wordCount struct {
		word  string
		count int
	}

	words := make([]wordCount, 0, len(results))
	for word, count := range results {
		words = append(words, wordCount{word, count})
	}

	sort.Slice(words, func(i, j int) bool {
		if words[i].count == words[j].count {
			return words[i].word < words[j].word
		}
		return words[i].count > words[j].count
	})

	fmt.Println("Top 30 most frequent words:")
	limit := 30
	if len(words) < limit {
		limit = len(words)
	}
	for i := 0; i < limit; i++ {
		fmt.Printf(" - \"%s\": %d\n", words[i].word, words[i].count)
	}
	fmt.Printf("\nTotal unique words: %d\n", len(results))
}
