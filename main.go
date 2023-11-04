package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

type Data struct {
	Answer []string
}

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

func deDupe(a []string) []string {
	existList := make(map[string]bool)
	list := []string{}
	for _, i := range a {
		if _, v := existList[i]; !v {
			existList[i] = true
			list = append(list, i)
		}
	}
	return list
}

func findAnswers(query string) []string {
	answer := wordHash[query]
	l := len(query)
	if l > 3 {
		a := strings.Split(query, "")
		answer = append(answer, findAnswers(strings.Join(a[1:l], ""))...)
		for i := 1; i < l-1; i++ {
			b := make([]string, l-1)
			copy(b, a[0:i])
			b = append(b, a[i+1:l]...)
			answer = append(answer, findAnswers(strings.Join(b, ""))...)
		}
		answer = append(answer, findAnswers(strings.Join(a[0:l-1], ""))...)
	}
	return answer
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("search")
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, &Data{Answer: deDupe(findAnswers(hash(query)))})
}

// Set up a webserver
func main() {
	// Check environment variables
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	address := fmt.Sprintf("%s:%s", "0.0.0.0", httpPort)
	wordFile := os.Getenv("WORDFILE")
	if wordFile == "" {
		wordFile = "words"
	}
	wordHash = readWords(wordFile)
	http.HandleFunc("/", formHandler)
	log.Fatal(http.ListenAndServe(address, nil))
}
