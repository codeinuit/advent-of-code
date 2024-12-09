package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocationParser(t *testing.T) {
	lp := LocationParser{}

	lp.AddLine(3, 4)
	lp.AddLine(4, 3)
	lp.AddLine(2, 5)
	lp.AddLine(1, 3)
	lp.AddLine(3, 9)
	lp.AddLine(3, 3)

	assert.Equal(t, 11, lp.Result())
}
