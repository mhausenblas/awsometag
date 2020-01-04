package main

import (
	"log"
)

func tags3bucket(arn, tag string) error {
	k, v := compart(tag)
	log.Printf("Tagging S3 bucket %s with %s:%s", arn, k, v)
	return nil
}

func tagiamuser(arn, tag string) error {
	k, v := compart(tag)
	log.Printf("Tagging IAM user %s with %s:%s", arn, k, v)
	return nil
}
