// All networking services-related tagging functions
package main

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

// tagrds tags an RDS resource
func tagrds(arnres arn.ARN, key, value string) error {
	// resource types as per these docs:
	// https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_AddTagsToResource.html
	// for example:
	// arn:aws:rds:eu-west-1:123456789012:db:mydb
	region := arnres.Region
	dbname := strings.Split(arnres.Resource, ":")[1]
	log.Printf("Tagging RDS database '%s' in region '%s' with %s:%s",
		dbname, region, key, value)
	svc := rds.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion(region))
	_, err := svc.AddTagsToResource(&rds.AddTagsToResourceInput{
		ResourceName: aws.String(arnres.String()),
		Tags: []*rds.Tag{
			{
				Key:   aws.String(key),
				Value: aws.String(value),
			},
		},
	})
	return err
}
