package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/zkfmapf123/ml/tools"
)

func calcSigmoid(z float64) float64 {
	return 1 / (1 + math.Exp(-z))
}

func logisticRegression(data [][]float64, learningRate float64, iterations int) (float64, float64, float64, float64) {
	var w, b float64 = 0, 0
	n := float64(len(data))

	// ✅ 평균과 표준편차 계산 (Z-score 정규화용)
	sumX, sumX2 := 0.0, 0.0
	for _, row := range data {
		sumX += row[0]
		sumX2 += row[0] * row[0]
	}
	meanX := sumX / n
	stdX := math.Sqrt((sumX2 / n) - (meanX * meanX))

	for i := 0; i < iterations; i++ {
		var dw, db float64 = 0, 0

		for _, row := range data {
			x := (row[0] - meanX) / stdX // ✅ 표준화 적용
			y := row[1]

			predict := calcSigmoid(w*x + b)
			errorCount := predict - y

			dw += errorCount * x
			db += errorCount
		}

		w -= learningRate * (dw / n)
		b -= learningRate * (db / n)
	}

	fmt.Println("최종 w:", w, "최종 b:", b)
	return w, b, meanX, stdX
}

type LogisticResponseParams struct {
	Input        float64
	PredictRatio string
}

func Logistic(filename string, x float64) LogisticResponseParams {
	_data := tools.LoadCSV(fmt.Sprintf("./public/%s", filename))

	var data [][]float64
	for _, v := range _data {
		v1, _ := strconv.ParseFloat(v[0], 64)
		v2, _ := strconv.ParseFloat(v[1], 64)
		data = append(data, []float64{v1, v2})
	}

	// 학습 실행
	w, b, meanX, stdX := logisticRegression(data, 0.1, 10000) // 학습률 0.1, 반복 횟수 10000

	// ✅ 예측 시에도 표준화 적용
	normalizedX := (x - meanX) / stdX
	predict := calcSigmoid(w*normalizedX + b)
	fmt.Println("w*x + b:", w*normalizedX+b, "sigmoid:", predict)

	return LogisticResponseParams{
		Input:        x,
		PredictRatio: fmt.Sprintf("%.2f%%", predict*100),
	}
}
