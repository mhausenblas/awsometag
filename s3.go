// All S3 service-related tagging functions
package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// tags3object tags an S3 object in a bucket
func tags3object(region, bucketname, objectname, key, value string) error {
	log.Printf("Tagging S3 object '%s' in bucket '%s' with %s:%s", objectname, bucketname, key, value)
	svc := s3.New(session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(region),
		},
	})))
	_, err := svc.PutObjectTagging(&s3.PutObjectTaggingInput{
		Bucket: aws.String(bucketname),
		Key:    aws.String(objectname),
		Tagging: &s3.Tagging{
			TagSet: []*s3.Tag{
				{
					Key:   aws.String(key),
					Value: aws.String(value),
				},
			},
		},
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Println(aerr.Error())
			}
		} else {
			log.Println(err.Error())
		}
	}
	return err
}

// tags3bucket tags an S3 bucket
func tags3bucket(region, bucketname, key, value string) error {
	log.Printf("Tagging S3 bucket '%s' with %s:%s", bucketname, key, value)
	svc := s3.New(session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(region),
		},
	})))
	_, err := svc.PutBucketTagging(&s3.PutBucketTaggingInput{
		Bucket: aws.String(bucketname),
		Tagging: &s3.Tagging{
			TagSet: []*s3.Tag{
				{
					Key:   aws.String(key),
					Value: aws.String(value),
				},
			},
		},
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Println(aerr.Error())
			}
		} else {
			log.Println(err.Error())
		}
	}
	return err
}
