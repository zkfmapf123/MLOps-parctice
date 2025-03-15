package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_linear(t *testing.T) {

	model_1 := LinearCalc("data_1.csv", 5.0)
	model_2 := LinearCalc("data_1.csv", 1.732)
	assert.Equal(t, model_1.predictNum, "10")
	assert.Equal(t, model_2.predictNum, "3")
}

// 광고비를 얼마나 태우면 매출이 나올까?
func Test_linear2(t *testing.T) {

	a := LinearCalc("data_2.csv", 5)
	b := LinearCalc("data_2.csv", 13.5)

	// 광고비에 5만원을 태운다면 ?
	assert.Equal(t, a.predictNum, "58")

	// 광고비에 13만 5천원을 태운다면 ?
	assert.Equal(t, b.predictNum, "157")

}
