package mapred

import (
	"regexp"
	"strings"
	"sync"
)

type MapReduce struct {
}

// todo implement mapreduce

func (mr *MapReduce) wordCountMapper(text string) []KeyValue {
	// Filter out special characters and numerals
	re := regexp.MustCompile(`[^a-zA-Z\s]+`)
	cleanText := re.ReplaceAllString(text, " ")
	words := strings.Fields(cleanText) // Split by whitespace

	var result []KeyValue
	for _, word := range words {
		if word != "" {
			result = append(result, KeyValue{Key: strings.ToLower(word), Value: 1})
		}
	}

	return result
}

func (mr *MapReduce) wordCountReducer(key string, values []int) KeyValue {
	sum := 0
	for _, v := range values {
		sum += v
	}

	return KeyValue{Key: key, Value: sum}
}

func (mr *MapReduce) mapPhase(input []string) []KeyValue {
	mapResults := make(chan []KeyValue, len(input))
	var wg sync.WaitGroup

	for _, line := range input {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			mapResults <- mr.wordCountMapper(text)
		}(line)
	}

	// Close the channel when all mappers are done
	go func() {
		wg.Wait()
		close(mapResults)
	}()

	// Collect all mapped results
	var allMapped []KeyValue
	for kvList := range mapResults {
		allMapped = append(allMapped, kvList...)
	}

	return allMapped
}

func (mr *MapReduce) shufflePhase(mapped []KeyValue) map[string][]int {
	intermediate := make(map[string][]int)
	for _, kv := range mapped {
		intermediate[kv.Key] = append(intermediate[kv.Key], kv.Value)
	}
	return intermediate
}

func (mr *MapReduce) reducePhase(intermediate map[string][]int) map[string]int {
	resultChan := make(chan KeyValue, len(intermediate))
	var wg sync.WaitGroup

	for key, values := range intermediate {
		wg.Add(1)
		go func(k string, v []int) {
			defer wg.Done()
			resultChan <- mr.wordCountReducer(k, v)
		}(key, values)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	finalResults := make(map[string]int)
	for kv := range resultChan {
		finalResults[kv.Key] = kv.Value
	}

	return finalResults
}

func (mr *MapReduce) Run(input []string) map[string]int {
	mapped := mr.mapPhase(input)
	shuffled := mr.shufflePhase(mapped)
	results := mr.reducePhase(shuffled)

	return results
}
