package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// UploadToS3 to upload a file to s3
func UploadToS3(uploadInput *s3manager.UploadInput) error {
	awsSession := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(awsSession)
	uploadOutput, err := uploader.Upload(uploadInput)
	if err != nil {
		return fmt.Errorf("File to upload to S3, %v", err)
	}
	log.Println("Succesfully uploaded to S3", uploadOutput.Location)
	return nil
}
