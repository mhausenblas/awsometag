// All networking services-related tagging functions
package main

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// tagsecret tags a secret maintained by the Secrets Manager
func tagsecret(arnres arn.ARN, key, value string) error {
	// resource types as per these docs:
	// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_TagResource.html
	// for example:
	// arn:aws:secretsmanager:us-west-2:123456789102:secret:mysecret-123456
	region := arnres.Region
	secretname := strings.Split(strings.Split(arnres.Resource, ":")[1], "-")[0]
	log.Printf("Tagging secret '%s' in region '%s' with %s:%s",
		secretname, region, key, value)
	svc := secretsmanager.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion(region))
	_, err := svc.TagResource(&secretsmanager.TagResourceInput{
		SecretId: aws.String(secretname),
		Tags: []*secretsmanager.Tag{
			{
				Key:   aws.String(key),
				Value: aws.String(value),
			},
		},
	})
	return err
}
