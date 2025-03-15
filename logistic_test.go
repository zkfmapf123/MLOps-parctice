package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sigmoid(t *testing.T) {

	answers := []string{
		"0.00005",
		"0.00669",
		"0.26894",
	}

	for i, v := range []float64{-10, -5, -1} {

		sig := calcSigmoid(v)

		assert.Equal(t, fmt.Sprintf("%.5f", sig), answers[i])
	}
}
