package s3

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetRegion(t *testing.T) {
	expected := os.Getenv("TEST_BUCKET_REGION")
	actual, err := GetRegion(os.Getenv("TEST_BUCKET"))
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
