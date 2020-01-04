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
		log.Fatalln("Need resource ARN and tags")
	}
	// the ARN of the resource to tag is supposed to be the first argument:
	res := os.Args[1]
	// a comma-separated list of tags os supposed to be the second argument:
	tags := os.Args[2]
	// first try to guess the resource/service type:
	rtype, err := guesstype(res)
	if err != nil {
		log.Fatalf("Can't guess the type of resource based on ARN %s", res)
	}
	// and finally try to tag the resource/service with the tags provided:
	rtag(res, rtype, tags)
}

// guesstype extracts the resource/service type of the ARN, see also:
// https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
func guesstype(res string) (string, error) {
	arnres, err := arn.Parse(res)
	if err != nil {
		return "", err
	}
	return arnres.Service, nil
}

// rtag tags a given resource with a certain type with a comma-separated
// list of tags or fails if it doesn't support the resource type
func rtag(res, rtype, tags string) (err error) {
	taglist := expand(tags)
	for _, tag := range taglist {
		switch rtype {
		case "iam":
			err = tagiamuser(res, tag)
		case "s3":
			err = tags3bucket(res, tag)
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
