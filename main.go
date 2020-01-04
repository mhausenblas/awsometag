package main

import (
	"fmt"
	"log"
	"os"

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
	rtype, err := guesstype(res)
	if err != nil {
		log.Fatalf("Can't guess the type of resource based on ARN %s", res)
	}
	log.Printf("Tagging %s of type %s with %s", res, rtype, tags)
	rtag(res, rtype)
}

func guesstype(res string) (string, error) {
	// as per https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	arnres, err := arn.Parse(res)
	if err != nil {
		return "", err
	}
	return arnres.Service, nil
}

func rtag(res, rtype string) error {
	switch rtype {
	case "iam":
		fmt.Println("TAG IAM")
	case "s3":
		fmt.Println("TAG S3")
	default:
		return fmt.Errorf("Don't know how to tag resources of type %s", rtype)
	}
	return nil
}
