// All container services-related tagging functions
package main

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/eks"
)

// tageks tags an EKS cluster or an EKS managed node group
func tageks(region, arns, key, value string) error {
	switch {
	case strings.Contains(arns, "cluster/"): // arn:aws:eks:us-west-2:123456789102:cluster/somecluster
		clustername := strings.Split(arns, "/")[1]
		log.Printf("Tagging EKS cluster '%s' in region '%s' with %s:%s",
			clustername, region, key, value)
	case strings.Contains(arns, "nodegroup/"): // arn:aws:eks:us-west-2:123456789102:nodegroup/somecluster/ng-2fb16d2a/123456789123456789
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

// tagecr tags an ECR repository
func tagecr(region, arns, key, value string) error {
	// arn:aws:ecr:us-west-2:123456789102:repository/somerepo
	reponame := strings.Split(arns, "/")[1]
	log.Printf("Tagging ECR repository '%s' in region '%s' with %s:%s",
		reponame, region, key, value)

	svc := ecr.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion(region))
	_, err := svc.TagResource(&ecr.TagResourceInput{
		ResourceArn: aws.String(arns),
		Tags: []*ecr.Tag{
			{
				Key:   aws.String(key),
				Value: aws.String(value),
			},
		},
	})
	return err
}

// tagecs tags the following ECS resources: capacity providers, clusters, tasks,
// task definitions, services, and container instances.
func tagecs(region, arns, key, value string) error {
	switch {
	case strings.Contains(arns, "capacity-provider/"): // arn:aws:ecs:us-west-2:123456789102:capacity-provider/FARGATE_SPOT
		capacityprovidername := strings.Split(arns, "/")[1]
		log.Printf("Tagging ECS capacity provider '%s' in region '%s' with %s:%s", capacityprovidername, region, key, value)
	case strings.Contains(arns, "cluster/"): // arn:aws:ecs:us-west-2:123456789102:cluster/somecluster
		clustername := strings.Split(arns, "/")[1]
		log.Printf("Tagging ECS cluster '%s' in region '%s' with %s:%s",
			clustername, region, key, value)
	case strings.Contains(arns, "task/"): // arn:aws:ecs:us-west-2:123456789102:task/somecluster/123456789123456789
		clustername := strings.Split(arns, "/")[1]
		taskname := strings.Split(arns, "/")[2]
		log.Printf("Tagging ECS task '%s' in cluster '%s' in region '%s' with %s:%s",
			taskname, clustername, region, key, value)
	case strings.Contains(arns, "task-definition/"): // arn:aws:ecs:us-west-2:123456789102:task-definition/sometd:42
		taskdefinitionname := strings.Split(arns, "/")[1]
		log.Printf("Tagging ECS task definition '%s' in region '%s' with %s:%s",
			taskdefinitionname, region, key, value)
	case strings.Contains(arns, "service/"): // arn:aws:ecs:us-west-2:123456789102:service/aservice
		servicenname := strings.Split(arns, "/")[1]
		log.Printf("Tagging ECS service '%s' in region '%s' with %s:%s",
			servicenname, region, key, value)
	case strings.Contains(arns, "container-instance/"): // arn:aws:ecs:us-west-2:123456789102:container-instance/ancinstance1234
		containerinstancenname := strings.Split(arns, "/")[1]
		log.Printf("Tagging ECS container instance '%s' in region '%s' with %s:%s",
			containerinstancenname, region, key, value)
	}
	svc := ecs.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion(region))
	_, err := svc.TagResource(&ecs.TagResourceInput{
		ResourceArn: aws.String(arns),
		Tags: []*ecs.Tag{
			{
				Key:   aws.String(key),
				Value: aws.String(value),
			},
		},
	})
	return err
}
