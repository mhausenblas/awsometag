// All networking services-related tagging functions
package main

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
)

// taglb tags a load balancer (CLB, ALB, NLB)
func taglb(arnres arn.ARN, key, value string) error {
	// arn:aws:elasticloadbalancing:us-west-2:123456789102:loadbalancer/app/my-test-alb/1234567890
	region := arnres.Region
	lbtype := strings.Split(arnres.Resource, "/")[1]
	lbname := strings.Split(arnres.Resource, "/")[2]
	log.Printf("Tagging load balancer '%s' of type '%s' in region '%s' with %s:%s",
		lbname, lbtype, region, key, value)
	svc := elb.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion(region))
	_, err := svc.AddTags(&elb.AddTagsInput{
		LoadBalancerNames: []*string{
			aws.String(lbname),
		},
		Tags: []*elb.Tag{
			{
				Key:   aws.String(key),
				Value: aws.String(value),
			},
		},
	})
	return err
}
