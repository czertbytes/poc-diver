package main

import (
	"bytes"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const OutputBucket = "poc-diver-output"

func HandleRequest(ctx context.Context, s3Events events.S3Event) (string, error) {
	s, err := session.NewSession()
	if err != nil {
		return `{"result":"creating aws session failed"}`, err
	}

	for _, record := range s3Events.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key

		input, err := s3.New(s).GetObject(&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})
		if err != nil {
			return `{"result":"reading input file failed"}`, err
		}
		defer input.Body.Close()

		diver := NewDiver(
			WithConstraint("DATA_TYPE_BOP", NewNumericConstraint[int8](0, 5, 2)),
		)

		var output bytes.Buffer
		if err := diver.Run(input.Body, &output); err != nil {
			return `{"result":"running diver failed"}`, err
		}

		if _, err = s3.New(s).PutObject(&s3.PutObjectInput{
			Bucket:      aws.String(OutputBucket),
			Key:         aws.String(key),
			ACL:         aws.String("private"),
			Body:        bytes.NewReader(output.Bytes()),
			ContentType: aws.String("text/csv"),
		}); err != nil {
			return `{"result":"storing output file failed"}`, err
		}
	}

	return `{"result":"done"}`, nil
}

func main() {
	lambda.Start(HandleRequest)
}
