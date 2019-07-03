package middleware

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func customSortingTest(t *testing.T) {
	initialQueue := []*Queue{
		{
			Domain:   "alpha",
			Weight:   4,
			Priority: 1,
		},
		{
			Domain:   "omega",
			Weight:   2,
			Priority: 2,
		},
		{
			Domain:   "beta",
			Weight:   5,
			Priority: 5,
		},
	}
	expectedQueue := []*Queue{
		{
			Domain:   "alpha",
			Weight:   5,
			Priority: 5,
		},
		{
			Domain:   "omega",
			Weight:   4,
			Priority: 1,
		},
		{
			Domain:   "beta",
			Weight:   2,
			Priority: 2,
		},
	}
	fmt.Println("test sort")
	sorted := CustomSorting(initialQueue)
	assert.Equal(t, sorted, expectedQueue)
}
