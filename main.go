package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws/arn"
)

func main() {
	if len(os.Args[1:]) < 2 {
		log.Fatalln("Need both the resource ARN and tags, sorry :(")
	}
	// the ARN of the resource to tag is supposed to be the first argument:
	arns := os.Args[1]
	// a comma-separated list of tags os supposed to be the second argument:
	tags := os.Args[2]
	// first try to guess the resource type:
	rtype, err := guesstype(arns)
	if err != nil {
		log.Fatalf("Can't guess the type of resource based on the ARN %s", arns)
	}
	// and finally try to tag the resource/with the tags provided:
	err = rtag(arns, rtype, tags)
	if err != nil {
		log.Fatalln(err)
	}
}

// rtag tags a resource with ARN arns and a certain type rtype
// with a comma-separated list of tags or fails if it
// doesn't support the resource type
func rtag(arns, rtype, tags string) (err error) {
	taglist := expand(tags)
	for _, tag := range taglist {
		arnres, key, value, err := preflight(arns, tag)
		if err != nil {
			return err
		}
		region := arnres.Region
		switch rtype {
		case "iam":
			iamtype := strings.Split(arnres.Resource, "/")[0]
			switch iamtype {
			case "user": // arn:aws:iam::123456789102:user/xxx
				username := strings.Split(arnres.Resource, "/")[1]
				err = tagiamuser(username, key, value)
			case "role": // arn:aws:iam::123456789102:role/xxx
				rolename := strings.Split(arnres.Resource, "/")[1]
				err = tagiamrole(rolename, key, value)
			default:
				return fmt.Errorf("I know how to tag IAM users and roles, and %s seems to be neither", arns)
			}
		case "s3":
			if r := os.Getenv("S3_ENDPOINT_REGION"); r != "" {
				region = r
			}
			// note that the following is a simplified case distinction since there
			// are other resource types (accesspoint and jobs) defined by the S3 service
			// see https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazons3.html#amazons3-resources-for-iam-policies
			switch {
			case strings.Contains(arnres.Resource, "/"): // arn:aws:s3:::abucket/anobject
				bucketname := strings.Split(arnres.Resource, "/")[0]
				objectname := strings.Split(arnres.Resource, "/")[1]
				err = tags3object(region, bucketname, objectname, key, value)
			case !strings.Contains(arnres.Resource, "/"): // arn:aws:s3:::abucket
				bucketname := arnres.Resource
				err = tags3bucket(region, bucketname, key, value)
			default:
				return fmt.Errorf("I know how to tag S3 buckets and objects, and %s seems to be neither", arns)
			}
		case "eks":
			switch {
			case strings.HasPrefix(arnres.Resource, "cluster"), // arn:aws:eks:*:*:cluster
				strings.HasPrefix(arnres.Resource, "nodegroup"): // arn:aws:eks:*:*:nodegroup
				err = tageks(region, arns, key, value)
			default:
				return fmt.Errorf("I know how to tag EKS clusters and managed node groups, and %s seems to be neither", arns)
			}
		case "ecr":
			switch {
			case strings.HasPrefix(arnres.Resource, "repository"): // arn:aws:ecr:*:*:repository
				err = tagecr(region, arns, key, value)
			default:
				return fmt.Errorf("I know how to tag ECR repos, and %s seems not to be one", arns)
			}
		case "ecs":
			switch {
			case strings.HasPrefix(arnres.Resource, "capacity-provider"), // arn:aws:ecs:*:*:capacity-provider/*
				strings.HasPrefix(arnres.Resource, "cluster"),           // arn:aws:ecs:*:*:cluster/*
				strings.HasPrefix(arnres.Resource, "task"),              // arn:aws:ecs:*:*:task/cluster/*
				strings.HasPrefix(arnres.Resource, "task-definition"),   // arn:aws:ecs:*:*:task-definition/*
				strings.HasPrefix(arnres.Resource, "service"),           // arn:aws:ecs:*:*:service/*
				strings.HasPrefix(arnres.Resource, "capacity-provider"): // arn:aws:ecs:*:*:capacity-provider/*
				err = tagecs(region, arns, key, value)
			default:
				return fmt.Errorf("I only know how to tag ECS capacity providers, clusters, tasks, task definitions, services, and container instances, and %s seems to be not one of those", arns)
			}
		default:
			return fmt.Errorf("Don't know how to tag resources of type %s", rtype)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// guesstype extracts the resource type of the ARN, see also:
// https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
func guesstype(arns string) (string, error) {
	arnres, err := arn.Parse(arns)
	if err != nil {
		return "", err
	}
	return arnres.Service, nil
}

// expand splits a tag of the form 'key1=val1, key2=val2' into a string slice
func expand(tags string) []string {
	raw := strings.Split(tags, ",")
	clean := []string{}
	for _, tag := range raw {
		clean = append(clean, strings.TrimSpace(tag))
	}
	return clean
}

// compart splits a tag of the form key=val into its components and tries to
// provide sensible values if it fails to do so
func compart(tag string) (key, val string) {
	if !strings.Contains(tag, "=") {
		return "", ""
	}
	kv := strings.Split(tag, "=")
	if len(kv) == 2 {
		return kv[0], kv[1]
	}
	return kv[0], ""
}

// preflight converts the ARN string and a tag into an ARN object and a
// key-value pair; a convenience function, only.
func preflight(arnres, tag string) (a arn.ARN, k string, v string, err error) {
	k, v = compart(tag)
	a, err = arn.Parse(arnres)
	if err != nil {
		return a, "", "", err
	}
	return
}
