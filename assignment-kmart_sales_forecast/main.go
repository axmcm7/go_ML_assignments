package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sajari/regression"
	"github.com/tobgu/qframe"
)

// checks if an error is nil, if not--prints the error
func errIsNil(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

// converts the json files into a Qframe dataframe
func readInJSONData(fileName string) qframe.QFrame {
	jsonFile, err := os.Open(fileName)
	errIsNil(err)

	// Parse the file
	r := qframe.ReadJSON(jsonFile)
	return r
}

// sets up the regression model object before the regression is run
func initializeRegressionData(reg regression.Regression, df qframe.QFrame) regression.Regression {
	reg.SetObserved("Sales")

	cols := df.ColumnNames()
	var xView []qframe.FloatView
	yView, err := df.FloatView("Sales")
	errIsNil(err)

	for i := 0; i < len(cols); i++ {
		// skip the target column
		if cols[i] == "Sales" {
			continue
		}
		colFloatView, err2 := df.FloatView(cols[i])
		errIsNil(err2)
		xView = append(xView, colFloatView)
	}

	for i := 0; i < df.Len(); i++ {
		dataForPointI := make([]float64, len(xView))
		for j := 0; j < len(xView); j++ {
			dataForPointI[j] = xView[j].ItemAt(i)
		}

		point := regression.DataPoint(yView.ItemAt(i), dataForPointI)
		reg.Train(point)

	}
	return reg
}

// makes predictions for each data point in the test set, returns one large string thereof
func generatePredictions(reg regression.Regression, df qframe.QFrame) string {
	var X []qframe.FloatView
	cols := df.ColumnNames()
	for i := 0; i < len(cols); i++ {
		colFloatView, _ := df.FloatView(cols[i])
		X = append(X, colFloatView)
	}

	var preds bytes.Buffer

	preds.WriteString("R2 score: " + fmt.Sprintf("%f", reg.R2) + "\n\n")
	preds.WriteString("Predictions for 'Item_Outlet_Sales':\n")

	for i := 0; i < df.Len(); i++ {
		var data []float64
		for j := 0; j < len(X); j++ {
			data = append(data, X[j].ItemAt(i))
		}

		pred, err := reg.Predict(data)
		errIsNil(err)

		preds.WriteString(fmt.Sprintf("%f", pred) + "\n")
	}

	return preds.String()
}

// saves the predictions string
func saveToFile(output string, savePath string) {
	err := ioutil.WriteFile(savePath, []byte(output), 0666)
	errIsNil(err)
}

func main() {
	trainDF := readInJSONData("./train.json")
	testDF := readInJSONData("./test.json")

	reg := new(regression.Regression)
	// load in the data
	regModel := initializeRegressionData(*reg, trainDF)

	// run the regression and print results
	regModel.Run()
	fmt.Printf("Regression formula:\n%v\n", regModel.Formula)

	// generate the predictions for the test data and save it
	predsString := generatePredictions(regModel, testDF)
	saveToFile(predsString, "./kmart_sales_predictions")
}
