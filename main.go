package main

import (
	"flag"
	"pi-square-io/go-aws-tags/functions"
)

var (
	region   = flag.String("region", "us-east-1", "your AWS region")
	resource = flag.String("resource", "ec2", "your resource")
	id       = flag.String("id", "id", "id,arn or name of resource")
	tagkey   = flag.String("tagkey", "tagkey", "key of tag resource")
	tagvalue = flag.String("tagvalue", "tagvalue", "value of tag resource")
)

func main() {

	flag.Parse()

	if *resource == "rds" {
		functions.TagRds(*id, *tagkey, *tagvalue, *region)
	}

	if *resource == "ec2" {
		var instanceId = []string{*id}
		functions.TagEc2(instanceId, *tagkey, *tagvalue, *region)
	}

	if *resource == "s3" {
		functions.TagS3(*id, *tagkey, *tagvalue, *region)
	}

	if *resource == "elb" {
		var albArns = []string{*id}
		functions.TagElb(albArns, *tagkey, *tagvalue, *region)
	}

	if *resource == "lambda" {
		functions.TagLambda(*id, *tagkey, *tagvalue, *region)
	}

	if *resource == "efs" {
		functions.TagEfs(*id, *tagkey, *tagvalue, *region)
	}
}
