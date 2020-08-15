package s3_utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func getTestBucket() string {
	bucket := os.Getenv("TEST_BUCKET")
	if len(bucket) == 0 {
		panic("can't get TEST_BUCKET from env")
	}
	return bucket
}

func getTestBucketRegion() string {
	region := os.Getenv("TEST_BUCKET_REGION")
	if len(region) == 0 {
		panic("can't get TEST_BUCKET_REGION from env")
	}
	return region
}

func getTestKey() string {
	key := os.Getenv("TEST_KEY")
	if len(key) == 0 {
		panic("can't get TEST_KEY from env")
	}
	return key
}

func TestGetRegion(t *testing.T) {
	expected := getTestBucketRegion()
	actual, err := GetRegion(getTestBucket())
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestListObjectsInBucket(t *testing.T) {
	actual, err := ListObjectsInBucket(getTestBucket(), getTestBucketRegion())
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
}

func TestGetObjectSinglePart(t *testing.T) {
	actual, err := GetObjectSinglePart(getTestBucket(), getTestKey(), getTestBucketRegion())
	assert.NoError(t, err)
	buf := make([]byte, *actual.ContentLength)
	_, err = actual.Body.Read(buf)
	assert.NoError(t, err)
	strr := string(buf)
	assert.NotEmpty(t, strr)
	assert.NotEmpty(t, buf)
}

func TestSessionWithGivenRegion(t *testing.T) {
	expectedRegion := "us-east-1"
	actual, err := sessionWithGivenRegion(expectedRegion, globalSess)
	assert.NoError(t, err)
	assert.Equal(t, expectedRegion, *actual.Config.Region)
}

func TestHeadObject(t *testing.T) {
	actual, err := HeadObject(getTestBucket(), getTestKey(), getTestBucketRegion())
	assert.NoError(t, err)
	assert.NotEmpty(t, actual.ETag)
}