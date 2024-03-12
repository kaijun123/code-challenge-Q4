package Q4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sum_to_n_a(t *testing.T) {
	result := sum_to_n_a(5)
	assert.Equal(t, 15, result)
}

func Test_sum_to_n_b(t *testing.T) {
	result := sum_to_n_b(5)
	assert.Equal(t, 15, result)
}

func Test_sum_to_n_c(t *testing.T) {
	result := sum_to_n_c(5)
	assert.Equal(t, 15, result)
}
