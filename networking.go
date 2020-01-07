// All networking services-related tagging functions
package main

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

// taglb tags a load balancer of type ALB or NLB
// note this is the v2 of the ELB API, see also:
// https://docs.aws.amazon.com/cli/latest/reference/elbv2/
func taglb(arnres arn.ARN, key, value string) error {
	// resource types as per these docs:
	// https://docs.aws.amazon.com/IAM/latest/UserGuide/list_elasticloadbalancingv2.html#elasticloadbalancingv2-resources-for-iam-policies
	// for example:
	// arn:aws:elasticloadbalancing:us-west-2:123456789102:loadbalancer/app/my-test-alb/1234567890
	region := arnres.Region
	lbtype := strings.Split(arnres.Resource, "/")[1]
	lbname := strings.Split(arnres.Resource, "/")[2]
	log.Printf("Tagging load balancer '%s' of type '%s' in region '%s' with %s:%s",
		lbname, lbtype, region, key, value)
	svc := elbv2.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion(region))
	_, err := svc.AddTags(&elbv2.AddTagsInput{
		ResourceArns: []*string{
			aws.String(arnres.String()),
		},
		Tags: []*elbv2.Tag{
			{
				Key:   aws.String(key),
				Value: aws.String(value),
			},
		},
	})
	return err
}

// taglb tags a load balancer of type CLB
// note this is the original ELB API, see also:
// https://docs.aws.amazon.com/cli/latest/reference/elb/
func taglbclassic(arnres arn.ARN, key, value string) error {
	// resource types as per these docs:
	// https://docs.aws.amazon.com/IAM/latest/UserGuide/list_elasticloadbalancing.html#elasticloadbalancing-resources-for-iam-policies
	// for example:
	// arn:aws:elasticloadbalancing:us-west-2:123456789102:loadbalancer/my-test-clb
	region := arnres.Region
	lbname := strings.Split(arnres.Resource, "/")[1]
	log.Printf("Tagging classic load balancer '%s' in region '%s' with %s:%s",
		lbname, region, key, value)
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
