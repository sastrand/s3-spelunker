package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const corpus0 = "cababbac"
const substring0 = "abba"
const corpus1 = "Only connect! That was the whole of her sermon. Only connect the prose and the passion, and both will be exalted, and human love will be seen at its height. Live in fragments no longer. Only connect and the beast and the monk, robbed of the isolation that is life to either, will die."
const substring1 = "Only connect and the"

func TestNaiveSubstringSearch_match(t *testing.T) {
	match, startIndex := NaiveSubstringSearch(corpus1, substring1)
	assert.True(t, match)
	assert.NotZero(t, startIndex)
}

func TestNaiveSubstringSearch_noMatch(t *testing.T) {
	match, startIndex := NaiveSubstringSearch(corpus1, substring0)
	assert.False(t, match)
	assert.Zero(t, startIndex)
}

func TestNaiveSubstringSearch_substringTooBig(t *testing.T) {
	match, startIndex := NaiveSubstringSearch(corpus0, substring1)
	assert.False(t, match)
	assert.Zero(t, startIndex)
}