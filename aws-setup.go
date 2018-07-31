package main

import (
	"context"

	"github.com/google/go-cloud/blob/s3blob"

	"github.com/google/go-cloud/blob"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	DefaultRegion = "ap-southeast-2"
)

func SetupAws(ctx context.Context, bucketName string) (*blob.Bucket, error) {
	awsConfig := &aws.Config{
		// Either hard-code the region or use AWS_REGION.
		Region: aws.String(DefaultRegion),
		// credentials.NewEnvCredentials assumes two environment variables are
		// present:
		// 1. AWS_ACCESS_KEY_ID, and
		// 2. AWS_SECRET_ACCESS_KEY.
		// If not just dont supply and it uses your machine local aws configured entries
		Credentials: credentials.NewEnvCredentials(),
	}
	awsSession := session.Must(session.NewSession(awsConfig))
	return s3blob.OpenBucket(ctx, awsSession, bucketName)
}

// UploadToS3 to upload a file to s3
//https://github.com/aws/aws-sdk-go#configuring-credentials
/*func UploadToS3(uploadInput *s3manager.UploadInput) error {
	log.Println("Uploading file to S3")
	awsSession := session.Must(session.NewSession(aws.NewConfig().WithMaxRetries(3).WithRegion(DefaultRegion)))
	uploader := s3manager.NewUploader(awsSession)
	uploadOutput, err := uploader.Upload(uploadInput)
	if err != nil {
		return fmt.Errorf("Failed to upload to S3, %v", err)
	}
	log.Println("Succesfully uploaded to S3", uploadOutput.Location)
	return nil
}*/
