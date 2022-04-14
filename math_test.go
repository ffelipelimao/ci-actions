package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sum(t *testing.T) {
	assert.Equal(t, Sum(10, 10), 20)
}
