package slices_test

import (
	"testing"
	"time"

	"github.com/doublecloud/transfer/library/go/slices"
	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	now := time.Now()

	origComplex := []time.Time{
		now,
		now.Add(1 * time.Second),
		now.Add(2 * time.Second),
		now.Add(3 * time.Second),
	}
	inputComplex := []time.Time{
		now,
		now.Add(1 * time.Second),
		now.Add(2 * time.Second),
		now.Add(3 * time.Second),
	}
	assert.NotEqual(t, origComplex, slices.Shuffle(inputComplex, nil))
}
