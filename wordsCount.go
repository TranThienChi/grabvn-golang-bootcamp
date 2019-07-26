package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/scanner"
)

// Slice filename paths for multiple text files in a folder
func SliceFilePaths(filePaths []string, folderDir string) []string {
	err := filepath.Walk(folderDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".txt" {
			return nil
		}

		filePaths = append(filePaths, path)
		return nil
	})

	if err != nil {
		panic(err)
	}
	return filePaths
}

// ConvertFileToLines converts text file to string lines
func ConvertFileToLines(file *os.File, stringLines chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringLines <- scanner.Text()
	}
}

// CountWord count the number of occurrences of a word
func CountWord(str string) map[string]int {
	result := make(map[string]int)
	var s scanner.Scanner
	s.Init(strings.NewReader(str))

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if tok == scanner.Ident {
			result[strings.ToLower(s.TokenText())] += 1
		}
	}
	return result
}

// HandleWork handle counting word in each string line
func HandleWork(stringLines chan string, wordCount chan map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for str := range stringLines {
		result := CountWord(str)
		fmt.Println(result)
		wordCount <- result
	}
}

// SumWordOccurrences summarize counting the number of occurrences of a word
func SumWordOccurrences(wordCount chan map[string]int, wordOccurrences chan map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := make(map[string]int)
	for res := range wordCount {
		for word, count := range res {
			result[word] += count
		}
	}
	wordOccurrences <- result
}

func main() {
	// Initialize local varialbles
	const maxWorker = 10
	wordOccurrences := make(chan map[string]int, 1)
	folderDir := "./multipleTextFiles"
	stringLines := make(chan string)
	wordCount := make(chan map[string]int)
	var fileGroup sync.WaitGroup
	var lineGroup sync.WaitGroup
	var sumGroup sync.WaitGroup

	// Slice filename paths for multiple text files in a folder
	filePaths := SliceFilePaths(make([]string, 0), folderDir)

	// Converts text file to string lines
	for _, file := range filePaths {
		file, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		fileGroup.Add(1)
		go ConvertFileToLines(file, stringLines, &fileGroup)
	}

	// Assign counting word for each worker
	lineGroup.Add(maxWorker)
	for i := 0; i < maxWorker; i++ {
		go HandleWork(stringLines, wordCount, &lineGroup)
	}

	// Summarize counting the number of occurrences of a word
	sumGroup.Add(1)
	go SumWordOccurrences(wordCount, wordOccurrences, &sumGroup)

	// Wait until finish counting word for all text files
	fileGroup.Wait()
	close(stringLines)

	lineGroup.Wait()
	close(wordCount)

	sumGroup.Wait()
	fmt.Printf("The number of occurrences of a word in a folder %q: %v", folderDir, <-wordOccurrences)
}
