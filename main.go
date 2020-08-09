package main

import (
    "context"
    "fmt"
    "os"
    
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Further Reading
// https://docs.aws.amazon.com/sdk-for-go/api/
// https://aws.amazon.com/blogs/developer/context-pattern-added-to-the-aws-sdk-for-go/

func main() {
    ctx := context.Background()
    sess := session.Must(session.NewSession())
    
    bucket := "sascha-test-data"
    region, err := s3manager.GetBucketRegion(ctx, sess, bucket, "us-west-2")
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "NotFound" {
             fmt.Fprintf(os.Stderr, "unable to find bucket %s's region not found\n", bucket)
        }
    }
    fmt.Printf("Bucket %s is in %s region\n", bucket, region)
}
