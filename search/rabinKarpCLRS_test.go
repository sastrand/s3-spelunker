package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var text0 = []uint64{1, 2, 3, 1, 2, 4, 9}
var pattern0 = []uint64{1, 2, 4}

func TestRabinKarpCLRS(t *testing.T) {
	match, startIndex := RabinKarpCLRS(text0, pattern0, 10)
	assert.True(t, match)
	assert.Equal(t, 3, startIndex)
}