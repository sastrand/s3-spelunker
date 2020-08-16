package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const corpus0 = "cabac"
const substring0 = "bac"
const corpus1 = "Only connect! That was the whole of her sermon. Only connect the prose and the passion, and both will be exalted, and human love will be seen at its height. Live in fragments no longer. Only connect and the beast and the monk, robbed of the isolation that is life to either, will die."
const substring1 = "Only connect and"

func TestNaiveSubstringSearch(t *testing.T) {
	//NaiveSubstringSearch(corpus0, substring0)
	actualMatch, actualStartIndex := NaiveSubstringSearch(corpus1, substring1)
	assert.True(t, actualMatch)
	assert.NotZero(t, actualStartIndex)
}