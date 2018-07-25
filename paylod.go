package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//PayloadCollection records
type PayloadCollection struct {
	Payloads []Payload `json:"users"`
}

// Payload format
type Payload struct {
	Name         string `json:"fullName"`
	Address      string `json:"address"`
	MobileNumber string `json:"mobile"`
}

const (
	S3StorageFolder = "paylod-upload"
	ACLPrivate      = "private"
	ContentType     = "application/octet-stream"
	DefaultRegion   = "ap-southeast-2"
)

//Upload the paylod to storage
func (payload *Payload) Upload() error {
	var fileName = fmt.Sprintf("%v%v", "test_", time.Now().UnixNano())
	log.Printf("File Name decided to be %v", fileName)
	b := new(bytes.Buffer)
	encodeErr := json.NewEncoder(b).Encode(payload)
	if encodeErr != nil {
		return encodeErr
	}
	uploadInput := s3manager.UploadInput{
		ACL:         aws.String(ACLPrivate),
		Bucket:      aws.String(S3StorageFolder),
		ContentType: aws.String(ContentType),
		Body:        b,
		Key:         aws.String(fileName),
	}
	return UploadToS3(&uploadInput)
}
