package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/zkfmapf123/ml/tools"
)

type DataPoint struct {
	CPU      float64
	Anomaly  int
	Distance float64
}

// 거리 계산
func euclideanDistance(x1, x2 float64) float64 {
	return math.Abs(x1 - x2)
}

func knnPredict(data []DataPoint, newCPU float64, k int) int {
	// 거리 계산
	for i := range data {
		data[i].Distance = euclideanDistance(data[i].CPU, newCPU)
	}

	// 거리순 정렬
	sort.Slice(data, func(i, j int) bool {
		return data[i].Distance < data[j].Distance
	})

	// K개의 이웃에서 이상(1)과 정상(0) 개수 확인
	count0, count1 := 0, 0
	for i := 0; i < k; i++ {
		if data[i].Anomaly == 0 {
			count0++
		} else {
			count1++
		}
	}

	// 다수결
	if count1 > count0 {
		return 1 // 이상 징후
	}
	return 0 // 정상
}

func knn(filename string, cpu int) string {

	data := tools.LoadCSV(fmt.Sprintf("./public/%s", filename))

	var dataPoint []DataPoint
	for _, v := range data {
		cpu, _ := strconv.ParseFloat(v[0], 64)
		anomaly, _ := strconv.ParseFloat(v[1], 64)

		dataPoint = append(dataPoint, DataPoint{CPU: cpu, Anomaly: int(anomaly)})
	}

	_newCpu := float64(cpu)

	/*
		k = 1 가장 가까운 1개만 참고 -> 노이즈에 민감
		k = 3 가장 가까운 3개 참고 -> 비교적 안정적
		k = 5 가장 가까운 5개 참고 -> 안정적이지만 결정경계가 부드러워짐
	*/
	result := knnPredict(dataPoint, _newCpu, 3)

	return fmt.Sprintf("CPU 사용량 %.1f일 때 이상 여부: %d (0=정상, 1=이상)\n", _newCpu, result)
}
