package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocationParser(t *testing.T) {
	//subject := "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"

	lp := LocationParser{}

	assert.Equal(t, 11, lp.Result())
}
