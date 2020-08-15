package s3Utils

import (
	"github.com/stretchr/testify/assert"
	"io"
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
	// Read returns EOF when the stream ends
	// https://tour.golang.org/methods/21
	if err != io.EOF {
		assert.NotEmpty(t, err)
	}
	assert.NotEmpty(t, buf)
}

func TestSessionWithGivenRegion(t *testing.T) {
	expectedRegion := "eu-south-1"
	if expectedRegion == *globalSess.Config.Region {
		panic("test requires expectedRegion different from default region")
	}
	actual, err := sessionWithGivenRegion(expectedRegion, globalSess)
	assert.NoError(t, err)
	assert.Equal(t, expectedRegion, *actual.Config.Region)
}

func TestHeadObject(t *testing.T) {
	actual, err := HeadObject(getTestBucket(), getTestKey(), getTestBucketRegion())
	assert.NoError(t, err)
	assert.NotEmpty(t, actual.ETag)
}

func TestGetObjectMultiPart(t *testing.T) {
	actual, err := GetObjectSinglePart(getTestBucket(), getTestKey(), getTestBucketRegion())
	assert.NoError(t, err)
	buf := make([]byte, *actual.ContentLength)
	_, err = actual.Body.Read(buf)
	if err != io.EOF {
		assert.NotEmpty(t, err)
	}
	assert.NotEmpty(t, buf)
}

func TestPutStringSinglePart(t *testing.T) {
	key := "test_upload.txt"
	body := "hello"
	err := PutStringSinglePart(getTestBucket(), key, getTestBucketRegion(), body)
	assert.NoError(t, err)
}

func TestPutStringMultiPart(t *testing.T) {
	key := "test_upload.txt"
	body := "hello"
	err := PutStringMultiPart(getTestBucket(), key, getTestBucketRegion(), body)
	assert.NoError(t, err)
}