package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws/arn"
)

func main() {
	if len(os.Args[1:]) < 3 {
		log.Fatalln("Need: resource ARN, region, and tags, sorry :(")
	}
	// the ARN of the resource to tag is supposed to be the first argument:
	arns := os.Args[1]
	// the region to apply the tagging in:
	region := os.Args[2]
	// a comma-separated list of tags os supposed to be the second argument:
	tags := os.Args[3]
	// first try to guess the resource type:
	rtype, err := guesstype(arns)
	if err != nil {
		log.Fatalf("Can't guess the type of resource based on the ARN %s", arns)
	}
	// and finally try to tag the resource/with the tags provided:
	rtag(region, arns, rtype, tags)
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

// rtag tags a resource with ARN arns and a certain type rtype
// with a comma-separated list of tags or fails if it
// doesn't support the resource type
func rtag(region, arns, rtype, tags string) (err error) {
	taglist := expand(tags)
	for _, tag := range taglist {
		arnres, key, value, err := preflight(arns, tag)
		if err != nil {
			return err
		}
		switch rtype {
		case "iam":
			iamtype := strings.Split(arnres.Resource, "/")[0]
			switch iamtype {
			case "user": // arn:aws:iam::123456789102:user/xxx
				username := strings.Split(arnres.Resource, "/")[1]
				err = tagiamuser(region, username, key, value)
			case "role": // arn:aws:iam::123456789102:role/xxx
				rolename := strings.Split(arnres.Resource, "/")[1]
				err = tagiamrole(region, rolename, key, value)
			default:
				return fmt.Errorf("Don't know how to tag resources of type %s", rtype)
			}
		case "s3":
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
				return fmt.Errorf("Don't know how to tag resources of type %s", rtype)
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
