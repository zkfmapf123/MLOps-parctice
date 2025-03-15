package linear

import (
	"fmt"
	"strconv"
	"strings"

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

	a := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	b := (sumY - a*sumX) / n
	return a, b
}

type ModelResponseParmas struct {
	x          float64
	predictNum string
}

func LinearCalc(filename string, x float64) ModelResponseParmas {

	rows := tools.LoadCSV(fmt.Sprintf("./public/%s", filename))

	var data []dataPointParams
	for i, row := range rows {
		if i == 0 {
			continue
		}

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
	// return fmt.Sprintf("Predict Y for X=%.2f Y=%2.f\n", x, predictY)
	return ModelResponseParmas{
		x:          x,
		predictNum: strings.Trim(fmt.Sprintf("%2.f", predictY), " "),
	}
}
