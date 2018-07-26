package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// UploadToS3 to upload a file to s3
//https://github.com/aws/aws-sdk-go#configuring-credentials
func UploadToS3(uploadInput *s3manager.UploadInput) error {
	log.Println("Uploading file to S3")
	awsSession := session.Must(session.NewSession(aws.NewConfig().WithMaxRetries(3).WithRegion(DefaultRegion)))
	uploader := s3manager.NewUploader(awsSession)
	uploadOutput, err := uploader.Upload(uploadInput)
	if err != nil {
		return fmt.Errorf("Failed to upload to S3, %v", err)
	}
	log.Println("Succesfully uploaded to S3", uploadOutput.Location)
	return nil
}
