package main

func readInCsvData()  {
	bs, err := ioutil.ReadFile("./NewsArticles_Top10Keywords.csv")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit()
	}
}

func main(args []string) {
	
}