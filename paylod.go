package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	S3StorageFolder = "test"
	ACLPrivate      = "private"
	ContentType     = "application/octet-stream"
)

//Upload the paylod to storage
func (payload *Payload) Upload() error {
	var fileName = fmt.Sprintf("%v/%v", "test_", time.Now().UnixNano())
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
	// return bucket.PutReader(storage_path, b, int64(b.Len()), contentType, acl, s3.Options{})
	return UploadToS3(&uploadInput)
}
