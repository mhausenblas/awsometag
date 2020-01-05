// All EKS service-related tagging functions
package main

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
)

// tageks tags an EKS cluster or an EKS managed node group
func tageks(region, arns, key, value string) error {
	switch {
	case strings.Contains(arns, "cluster/"): // arn:aws:eks:us-west-2:123456789102:cluster/somecluster
		clustername := strings.Split(arns, "/")[1]
		log.Printf("Tagging EKS cluster '%s' in region '%s' with %s:%s",
			clustername, region, key, value)
	case strings.Contains(arns, "nodegroup/"): // arn:aws:eks:us-west-2:123456789102:nodegroup/somecluster/ng-2fb16d2a/60b851gd-23eb-e15b-3234-34792affb491
		clustername := strings.Split(arns, "/")[1]
		nodegroupname := strings.Split(arns, "/")[2]
		log.Printf("Tagging managed node group '%s' in cluster '%s' in region '%s' with %s:%s",
			nodegroupname, clustername, region, key, value)
	}
	svc := eks.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion(region))
	_, err := svc.TagResource(&eks.TagResourceInput{
		ResourceArn: aws.String(arns),
		Tags:        map[string]*string{key: aws.String(value)},
	})
	return err
}
