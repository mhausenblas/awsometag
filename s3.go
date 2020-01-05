// All S3 service-related tagging functions
package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// tags3bucket tags an S3 bucket
func tags3bucket(region, arnres, tag string) error {
	res, k, v, err := preflight(arnres, tag)
	if err != nil {
		return err
	}
	log.Printf("Tagging S3 bucket '%s' with %s:%s", res.Resource, k, v)
	svc := s3.New(session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(region),
		},
	})))
	_, err = svc.PutBucketTagging(&s3.PutBucketTaggingInput{
		Bucket: aws.String(res.Resource),
		Tagging: &s3.Tagging{
			TagSet: []*s3.Tag{
				{
					Key:   aws.String(k),
					Value: aws.String(v),
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
