package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func tags3bucket(region, arnres, tag string) error {
	k, v := compart(tag)
	res, err := arn.Parse(arnres)
	if err != nil {
		return err
	}
	log.Printf("Tagging S3 bucket '%s' with %s:%s", res.Resource, k, v)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(region),
		},
	}))
	svc := s3.New(sess)
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
	return err
}

func tagiamuser(region, arn, tag string) error {
	k, v := compart(tag)
	log.Printf("Tagging IAM user %s with %s:%s", arn, k, v)
	return nil
}
