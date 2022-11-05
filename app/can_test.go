package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCellIsAlive(t *testing.T) {
	cell := NewCell()

	assert.True(t, cell.IsAlive())
	assert.True(t, true, "True is true!")
}

func TestNeigboursCounter(t *testing.T) {

	totalNeighbours := Neighbours()

	assert.True(t, totalNeighbours <= 8)
}
