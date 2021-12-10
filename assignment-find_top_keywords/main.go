package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	rake "github.com/afjoseph/RAKE.go"
	"github.com/tobgu/qframe"
)

// checks if an error is nil, if not--prints the error
func errIsNil(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

// converts the news.json file into a Qframe dataframe
func readInJsonData() qframe.QFrame {
	jsonFile, err := os.Open("./news.json")
	errIsNil(err)

	// Parse the file
	r := qframe.ReadJSON(jsonFile)
	return r
}

/*
concatenates all the title and content fields from the news data into
a string that can be fed into the RAKE algorithm
*/
func makeCorpusTextString(df qframe.QFrame) string {
	var text bytes.Buffer
	titles, _ := df.StringView("title")
	contents, _ := df.StringView("content")

	//for i := 0; i < 20; i++ {
	for i := 0; i < df.Len(); i++ {
		title := titles.ItemAt(i)
		content := contents.ItemAt(i)

		text.WriteString(*title)
		text.WriteString(*content)
	}

	return text.String()
}

// takes the RAKE alg output and formats it nicely into a string
func processRakeCandidates(candidates rake.PairList) string {
	var rakeOutput bytes.Buffer
	rakeOutput.WriteString(
		"KEYWORD PHRASE\t\t-->\t\tKEYWORD SCORE\n-----------------------------------------\n")

	for _, candidate := range candidates {
		candidateStr := fmt.Sprintf("%s --> %f\n", candidate.Key, candidate.Value)
		fmt.Println(candidateStr)
		//fmt.Printf("%s --> %f\n", candidate.Key, candidate.Value)
		rakeOutput.WriteString(candidateStr)
	}

	return rakeOutput.String()
}

// saves the formatted string of the RAKE output to the specified file path
func saveToFile(rakeOutput string, savePath string) {
	err := ioutil.WriteFile(savePath, []byte(rakeOutput), 0666)
	errIsNil(err)
}

func main() {
	numTopKeywordPhrases, err := strconv.Atoi(os.Args[1])
	errIsNil(err)

	df := readInJsonData()
	allTexts := makeCorpusTextString(df)

	// limit the output to the top K keyword phrases
	candidates := rake.RunRake(allTexts)[:numTopKeywordPhrases]
	//fmt.Printf("\nsize: %d\n", len(candidates))
	rakeStr := processRakeCandidates(candidates)

	saveToFile(rakeStr, fmt.Sprintf("./top%dKeywordPhrasesRAKE.txt", numTopKeywordPhrases))
}
