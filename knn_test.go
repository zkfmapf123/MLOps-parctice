package main

import (
	"fmt"
	"testing"
)

func Test_knn(t *testing.T) {

	v1 := knn("knn_data.csv", 35)
	fmt.Println(v1)

	v2 := knn("knn_data.csv", 60)
	fmt.Println(v2)
}
