package s3_utils

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// A session can be shared across parallel clients as long as it's not being modified.
// The most expensive part of creating a session is loading config values from the environment.
var globalSess *session.Session

func init() {
	globalSess = session.Must(session.NewSession())
}

// GetRegion : given a bucket, returns that bucket's region
func GetRegion(bucket string) (region string, err error) {
	return s3manager.GetBucketRegion(context.Background(), globalSess, bucket, "us-west-2")
}

func ListObjectsInBucket(bucket, region string) ([]*s3.Object, error) {
	svc := s3.New(globalSess, aws.NewConfig().WithRegion(region))
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	if err != nil {
		return nil, err
	}
	return resp.Contents, nil
}

func GetObjectSinglePart(bucket, key, region string) (*s3.GetObjectOutput, error) {
	svc := s3.New(globalSess, aws.NewConfig().WithRegion(region))
	return svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key: aws.String(key),
	})
}

// sessionWithRegion: if region in currentSession matches given region, returns currentRegion,
// otherwise creates a new session with given region
func sessionWithGivenRegion(region string, currentSession *session.Session) (*session.Session, error) {
	if *currentSession.Config.Region != region {
		return session.NewSession(&aws.Config{
			Region: aws.String(region),
		})
	}
	return currentSession, nil
}

// HeadObject: gets the metadata of an object, but the not the object itself
func HeadObject(bucket, key, region string) (*s3.HeadObjectOutput, error) {
	svc := s3.New(globalSess, aws.NewConfig().WithRegion(region))
	return svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key: aws.String(key),
	})
}

//// GetObjectMultiPart : downloads the specified object in parallel chunks
//// The tradeoffs to using this function are that it doesn't get the metadata in s3_utils.GetObjectOutput
//// that s3_utils.GetObject() returns in GetObjectSinglePart(). See headObject() for another way to get it.
//func GetObjectMultiPart(bucket, key, region string) ([]byte, error) {
//	sess, err := sessionWithGivenRegion(region, globalSess)
//	if err != nil {
//		return nil, err
//	}
//	downloader := s3manager.NewDownloader(sess)
//	buffer := aws.NewWriteAtBuffer([]byte{})
//	_, err = downloader.Download(buffer, &s3.GetObjectInput{
//		Bucket: aws.String(bucket),
//		Key: aws.String(key),
//	})
//	if err != nil {
//		return nil, err
//	}
//	return buffer.Bytes(), nil
//}
//
//func PutStringSinglePart(bucket, key, region, body string) error {
//	svc := s3.New(globalSess, aws.NewConfig().WithRegion(region))
//	_, err := svc.PutObject((&s3.PutObjectInput{}).
//		SetBucket(bucket).
//		SetKey(key).
//		SetBody(strings.NewReader(body)))
//	return err
//}
//
//func PutStringMultiPart(bucket, key, region, body string) error {
//	sess, err := sessionWithGivenRegion(region, globalSess)
//	if err != nil {
//		return err
//	}
//	uploader := s3manager.NewUploader(sess)
//	_, err = uploader.Upload(&s3manager.UploadInput{
//		Bucket: aws.String(bucket),
//		Key: aws.String(key),
//		Body: strings.NewReader(body),
//	})
//	return err
//}
