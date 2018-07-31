package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/go-cloud/blob"
	"github.com/pkg/errors"
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
	S3StorageBucketName = "paylod-upload"
	ACLPrivate          = "private"
	DefaultContentType  = "application/octet-stream"
	// DefaultRegion   = "ap-southeast-2"
)

//Upload the paylod to storage
func (payload *Payload) Upload(bucket *blob.Bucket, ctx context.Context) error {
	//Generate file name
	var fileName = fmt.Sprintf("%v%v.json", "paylod_", time.Now().UnixNano())
	log.Printf("File Name decided to be %v", fileName)
	//Encode payload object to JSON object
	b := new(bytes.Buffer)
	encodeErr := json.NewEncoder(b).Encode(payload)
	if encodeErr != nil {
		return errors.Wrap(encodeErr, "Unable to encode payload to JSON")
	}
	// Create a new bucket writter usng go-cloud
	//TODO: Cant find a way to supply ACL
	writerOptions := &blob.WriterOptions{ContentType: DefaultContentType}
	bucketWriter, bucketWriterErr := bucket.NewWriter(ctx, fileName, writerOptions)
	if bucketWriterErr != nil {
		return errors.Wrap(bucketWriterErr, "Error creating bucket writter")
	}
	//Write bytes to bucket using bucket writter
	_, writeErr := bucketWriter.Write(b.Bytes())
	if writeErr != nil {
		return errors.Wrap(writeErr, "Unable to write to bucket")
	}
	//Close the writter
	bucketWriterCloseErr := bucketWriter.Close()
	if bucketWriterCloseErr != nil {
		return errors.Wrap(bucketWriterCloseErr, "Error closing the bucket")
	}
	log.Printf("Succesfully uploaded %v to Cloud", fileName)
	return nil
}
