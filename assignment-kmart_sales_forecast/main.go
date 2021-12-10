package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
func readInJSONData(fileName string) qframe.QFrame {
	jsonFile, err := os.Open(fileName)
	errIsNil(err)

	// Parse the file
	r := qframe.ReadJSON(jsonFile)
	return r
}

// saves the formatted string of the RAKE output to the specified file path
func saveToFile(rakeOutput string, savePath string) {
	err := ioutil.WriteFile(savePath, []byte(rakeOutput), 0666)
	errIsNil(err)
}

func main() {
	trainDF := readInJSONData("./train.json")
	//testDF := readInJSONData("./test.json")

	fmt.Println(trainDF)
}
