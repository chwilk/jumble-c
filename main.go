package main

import (
	"net/http"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Vars that need to be global
var wordHash map[string][]string

// Hash function that alphabetizes letters in a word
func hash(s string) string {
	a := strings.Split(s, "")
	sort.Strings(a)
	return strings.Join(a, "")
}

// Parse dictionary into hash of slices
func readWords(wordFile string) map[string][]string {
	var wordHash = make(map[string][]string)
	wordList, err := os.Open(wordFile)
	if err != nil {
		log.Fatal(err)
	}
	defer wordList.Close()
	scanner := bufio.NewScanner(wordList)
	for scanner.Scan() {
		myWord := strings.Trim(scanner.Text(), " ")
		myHash := hash(myWord)
		wordHash[myHash] = append(wordHash[myHash], myWord)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return wordHash
}

func handler(w http.ResponseWriter, r *http.Request) {
	guess := "testing"
	fmt.Fprintf(w, "%s", wordHash[hash(guess)])
}

// Set up a webserver
func main() {
	wordHash = readWords("/words")
    http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}