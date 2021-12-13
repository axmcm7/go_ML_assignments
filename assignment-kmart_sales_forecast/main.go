package main

import (
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

func initializeDummyDataAndTrain(r regression.Regression) regression.Regression {
	r.SetObserved("Murders per annum per 1,000,000 inhabitants")
	r.SetVar(0, "Inhabitants")
	r.SetVar(1, "Percent with incomes below $5000")
	r.SetVar(2, "Percent unemployed")
	r.Train(
		regression.DataPoint(11.2, []float64{587000, 16.5, 6.2}),
		regression.DataPoint(13.4, []float64{643000, 20.5, 6.4}),
		regression.DataPoint(40.7, []float64{635000, 26.3, 9.3}),
		regression.DataPoint(5.3, []float64{692000, 16.5, 5.3}),
		regression.DataPoint(24.8, []float64{1248000, 19.2, 7.3}),
		regression.DataPoint(12.7, []float64{643000, 16.5, 5.9}),
		regression.DataPoint(20.9, []float64{1964000, 20.2, 6.4}),
		regression.DataPoint(35.7, []float64{1531000, 21.3, 7.6}),
		regression.DataPoint(8.7, []float64{713000, 17.2, 4.9}),
		regression.DataPoint(9.6, []float64{749000, 14.3, 6.4}),
		regression.DataPoint(14.5, []float64{7895000, 18.1, 6}),
		regression.DataPoint(26.9, []float64{762000, 23.1, 7.4}),
		regression.DataPoint(15.7, []float64{2793000, 19.1, 5.8}),
		regression.DataPoint(36.2, []float64{741000, 24.7, 8.6}),
		regression.DataPoint(18.1, []float64{625000, 18.6, 6.5}),
		regression.DataPoint(28.9, []float64{854000, 24.9, 8.3}),
		regression.DataPoint(14.9, []float64{716000, 17.9, 6.7}),
		regression.DataPoint(25.8, []float64{921000, 22.4, 8.6}),
		regression.DataPoint(21.7, []float64{595000, 20.2, 8.4}),
		regression.DataPoint(25.7, []float64{3353000, 16.9, 6.7}),
	)

	return r
}

func initializeDataAndTrainModel(reg regression.Regression, df qframe.QFrame) regression.Regression {
	reg.SetObserved("Item_Outlet_Sales")

	cols := df.ColumnNames()
	var X []qframe.FloatView
	y, err := df.FloatView("Item_Outlet_Sales")
	errIsNil(err)

	for i := 0; i < len(cols); i++ {
		// skip the target column
		if cols[i] == "Item_Outlet_Sales" {
			continue
		}

		reg.SetVar(i, cols[i])
		colFloatView, err2 := df.FloatView(cols[i])
		errIsNil(err2)
		X = append(X, colFloatView)
	}

	for i := 0; i < df.Len(); i++ {

		dataForPointI := make([]float64, len(X))
		for j := 0; j < len(X); j++ {
			dataForPointI[j] = X[j].ItemAt(i)
		}

		point := regression.DataPoint(y.ItemAt(i), dataForPointI)
		reg.Train(point)

		//if i == 1 {
		//	fmt.Println(y.ItemAt(i))
		//	fmt.Println(dataForPointI)
		//	break
		//}
	}

	// reg.Train()
	return reg
}

func generatePredictions(reg regression.Regression, df qframe.QFrame) {
	var X []qframe.FloatView
	cols := df.ColumnNames()
	for i := 0; i < len(cols); i++ {
		colFloatView, _ := df.FloatView(cols[i])
		X = append(X, colFloatView)
	}

	for i := 0; i < df.Len(); i++ {
		var data []float64
		for j := 0; j < len(X); j++ {
			data = append(data, X[j].ItemAt(i))
		}

		//fmt.Println(data)

		pred, err := reg.Predict(data)
		errIsNil(err)

		fmt.Println(pred)

		break
	}
}

func main() {
	trainDF := readInJSONData("./train.json")
	testDF := readInJSONData("./test.json")

	reg := new(regression.Regression)

	// load in the data and train the model
	//regModel := initializeDummyDataAndTrain(*reg)
	regModel := initializeDataAndTrainModel(*reg, trainDF)

	// run the regression and print results
	regModel.Run()
	//fmt.Printf("Regression formula:\n%v\n", regModel.Formula)
	fmt.Println(regModel.GetCoeffs())

	generatePredictions(regModel, testDF)
}
