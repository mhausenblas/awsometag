// All fundamental services-related tagging functions
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// tagiamrole tags an IAM role
func tagiamrole(rolename, key, value string) error {
	log.Printf("Tagging IAM role '%s' with %s:%s", rolename, key, value)
	svc := iam.New(session.New())
	_, err := svc.TagRole(&iam.TagRoleInput{
		RoleName: aws.String(rolename),
		Tags: []*iam.Tag{
			{
				Key:   aws.String(key),
				Value: aws.String(value),
			},
		},
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				log.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeLimitExceededException:
				log.Println(iam.ErrCodeLimitExceededException, aerr.Error())
			case iam.ErrCodeInvalidInputException:
				log.Println(iam.ErrCodeInvalidInputException, aerr.Error())
			case iam.ErrCodeConcurrentModificationException:
				log.Println(iam.ErrCodeConcurrentModificationException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				log.Println(iam.ErrCodeServiceFailureException, aerr.Error())
			default:
				log.Println(aerr.Error())
			}
		} else {
			log.Println(err.Error())
		}
	}
	return err
}

// tagiamuser tags an IAM user
func tagiamuser(username, key, value string) error {
	log.Printf("Tagging IAM user '%s' with %s:%s", username, key, value)
	svc := iam.New(session.New())
	_, err := svc.TagUser(&iam.TagUserInput{
		UserName: aws.String(username),
		Tags: []*iam.Tag{
			{
				Key:   aws.String(key),
				Value: aws.String(value),
			},
		},
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				log.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeLimitExceededException:
				log.Println(iam.ErrCodeLimitExceededException, aerr.Error())
			case iam.ErrCodeInvalidInputException:
				log.Println(iam.ErrCodeInvalidInputException, aerr.Error())
			case iam.ErrCodeConcurrentModificationException:
				log.Println(iam.ErrCodeConcurrentModificationException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				log.Println(iam.ErrCodeServiceFailureException, aerr.Error())
			default:
				log.Println(aerr.Error())
			}
		} else {
			log.Println(err.Error())
		}
	}
	return err
}

// tags3object tags an S3 object in a bucket
func tags3object(region, bucketname, objectname, key, value string) error {
	log.Printf("Tagging S3 object '%s' in bucket '%s' in region '%s' with %s:%s",
		objectname, bucketname, region, key, value)
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
	log.Printf("Tagging S3 bucket '%s' in region '%s' with %s:%s",
		bucketname, region, key, value)
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

// taglambda tags a Lambda function
func taglambda(arnres arn.ARN, key, value string) error {
	// arn:aws:lambda:us-west-2:123456789102:function:somefunc
	region := arnres.Region
	funcname := strings.Split(arnres.Resource, ":")[1]
	log.Printf("Tagging Lambda function '%s' in region '%s' with %s:%s",
		funcname, region, key, value)
	svc := lambda.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion(region))
	_, err := svc.TagResource(&lambda.TagResourceInput{
		Resource: aws.String(arnres.String()),
		Tags:     map[string]*string{key: aws.String(value)},
	})
	return err
}

// tagdynamodb tags a DynamoDB table
func tagdynamodb(arnres arn.ARN, key, value string) error {
	// arn:aws:dynamodb:us-west-2:123456789102:table/TheTable
	region := arnres.Region
	tablename := strings.Split(arnres.Resource, "/")[1]
	log.Printf("Tagging DynamoDB table '%s' in region '%s' with %s:%s",
		tablename, region, key, value)
	svc := dynamodb.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion(region))
	_, err := svc.TagResource(&dynamodb.TagResourceInput{
		ResourceArn: aws.String(arnres.String()),
		Tags: []*dynamodb.Tag{
			{
				Key:   aws.String(key),
				Value: aws.String(value),
			},
		},
	})
	return err
}

// tagec2 tags EC2 resources as defined in:
// https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazonec2.html#amazonec2-resources-for-iam-policies
func tagec2(arnres arn.ARN, key, value string) error {
	// arn:aws:ec2:us-west-2:123456789102:*/*
	region := arnres.Region
	resourcetype := strings.Split(arnres.Resource, "/")[0]
	resourceid := strings.Split(arnres.Resource, "/")[1]
	log.Printf("Tagging EC2 resource '%s' of type '%s' in region '%s' with %s:%s",
		resourceid, resourcetype, region, key, value)
	svc := ec2.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion(region))
	_, err := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{aws.String(resourceid)},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String(key),
				Value: aws.String(value),
			},
		},
	})
	return err
}

// tagsqs tags SQS queues as defined in:
// https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazonsqs.html#amazonsqs-resources-for-iam-policies
func tagsqs(arnres arn.ARN, key, value string) error {
	// arn:aws:sqs:us-west-2:123456789102:myqueue
	region := arnres.Region
	accountID := arnres.AccountID
	queuename := arnres.Resource
	log.Printf("Tagging SQS queue '%s' in region '%s' with %s:%s",
		queuename, region, key, value)
	// https://sqs.us-west-2.amazonaws.com/123456789102/myqueue
	queueURL := fmt.Sprintf("https://sqs.%s.amazonaws.com/%s/%s", region, accountID, queuename)
	svc := sqs.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion(region))
	_, err := svc.TagQueue(&sqs.TagQueueInput{
		QueueUrl: aws.String(queueURL),
		Tags:     map[string]*string{key: aws.String(value)},
	})
	return err
}
