package main

import (
	"flag"
	"pi-square-io/go-aws-tags/functions"
)

var (
	region   = flag.String("region", "us-east-1", "your AWS region")
	resource = flag.String("resource", "ec2", "your resource")
)

func main() {

	if *resource == "rds" {
		functions.TagRds("arn:aws:rds:us-east-1:974195321489:db:database-1", "Name", "test", *region)
	}

	if *resource == "ec2" {
		var instanceId = []string{"i-0ee77517f91d44f71"}
		functions.TagEc2(instanceId, "Name", "test", *region)
	}

	if *resource == "s3" {
		functions.TagS3("my-unique-435frzqrf", "Name", "test", *region)
	}

	if *resource == "alb" {
		var albArns = []string{"i-0ee77517f91d44f71"}
		functions.TagElb(albArns, "Name", "test", *region)
	}

	if *resource == "lambda" {
		functions.TagLambda("arn:aws:lambda:us-east-1:974195321489:function:test-tags", "Team", "DevOps", *region)
	}

	if *resource == "efs" {
		functions.TagEfs("fs-0fe29fa4a325c0807", "Dep", "DevOps", *region)
	}
}
