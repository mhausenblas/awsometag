package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/s3"
)

func preflight(arnres, tag string) (a arn.ARN, k string, v string, err error) {
	k, v = compart(tag)
	a, err = arn.Parse(arnres)
	if err != nil {
		return a, "", "", err
	}
	return
}

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
	return err
}

func tagiamrole(region, arnres, tag string) error {
	res, k, v, err := preflight(arnres, tag)
	if err != nil {
		return err
	}
	log.Printf("Tagging IAM role '%s' with %s:%s", res.Resource, k, v)
	svc := iam.New(session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(region),
		},
	})))
	_, err = svc.TagRole(&iam.TagRoleInput{
		RoleName: aws.String(res.Resource),
		Tags: []*iam.Tag{
			{
				Key:   aws.String(k),
				Value: aws.String(v),
			},
		},
	})
	return err
}

func tagiamuser(region, arnres, tag string) error {
	res, k, v, err := preflight(arnres, tag)
	if err != nil {
		return err
	}
	log.Printf("Tagging IAM user '%s' with %s:%s", res.Resource, k, v)
	svc := iam.New(session.New())
	r, err := svc.TagUser(&iam.TagUserInput{
		UserName: aws.String(res.Resource),
		Tags: []*iam.Tag{
			{
				Key:   aws.String(k),
				Value: aws.String(v),
			},
		},
	})
	log.Println(r)
	return err
}
