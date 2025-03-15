package linear

import (
	"fmt"
	"strconv"

	"github.com/zkfmapf123/ml/tools"
)

type dataPointParams struct {
	x float64
	y float64
}

// 선형회귀 알고리즘
func linearRegression(data []dataPointParams) (float64, float64) {

	// sumX : x 들의 합
	// sumY : y 들의 합
	// sumXY : x * y 들의 합
	// sumX2 : x^2 들의 합
	var sumX, sumY, sumXY, sumX2 float64

	n := float64(len(data))

	//
	for _, point := range data {
		sumX += point.x
		sumY += point.y
		sumXY = point.x * point.y
		sumX2 = point.x * point.x
	}

	a := (n*sumXY - sumY*sumY) / (n*sumX2 - sumX*sumX)
	b := (sumY - a*sumX) / n
	return a, b
}

func LinearCalc(x float64) {

	rows := tools.LoadCSV("./public/data_1.csv")

	var data []dataPointParams
	for _, row := range rows {

		r1, r2 := row[0], row[1]

		x, _ := strconv.ParseFloat(r1, 64)
		y, _ := strconv.ParseFloat(r2, 64)

		data = append(data, dataPointParams{x: x, y: y})
	}

	// y = ax + b
	// output : a
	// output : b
	a, b := linearRegression(data)

	// 기대값에 따른 예측값 생성
	predictY := a*x + b
	fmt.Printf("Predict Y for X=%.2f Y=%2.f\n", x, predictY)
}
