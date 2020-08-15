package s3

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// GetRegion : given a bucket, creates a new s3 session and returns bucket's region
// https://docs.aws.amazon.com/sdk-for-go/api/
// https://aws.amazon.com/blogs/developer/context-pattern-added-to-the-aws-sdk-for-go/
func GetRegion(bucket string) (region string, err error) {
	ctx := context.Background()
	sess := session.Must(session.NewSession())

	return s3manager.GetBucketRegion(ctx, sess, bucket, "us-west-2")
}
