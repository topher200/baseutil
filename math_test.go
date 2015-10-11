package baseutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandardDeviationInt(t *testing.T) {
	zeroStdDevList := []int{1, 1, 1, 1, 1}
	assert.Equal(t, 0.0, StandardDeviationInt(zeroStdDevList))

	knownStdDevList := []int{1, 2, 3, 4, 5}
	assert.InDelta(t, 1.5811, StandardDeviationInt(knownStdDevList), 0.0001)
}
