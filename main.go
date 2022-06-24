package main

import (
	"flag"
	"pi-square-io/go-aws-tags/functions"
)

var (
	region = flag.String("region", "us-east-1", "your AWS region")
)

func main() {

	functions.TagRds("arn:aws:rds:us-east-1:974195321489:db:database-1", "Name", "test", *region)

	var instanceId = []string{"i-0ee77517f91d44f71"}
	functions.TagEc2(instanceId, "Name", "test", *region)

	functions.TagS3("my-unique-435frzqrf", "Name", "test", *region)

	var albArns = []string{"i-0ee77517f91d44f71"}
	functions.TagElb(albArns, "Name", "test", *region)

	functions.TagLambda("arn:aws:lambda:us-east-1:974195321489:function:test-tags", "Team", "DevOps", *region)

	functions.TagEfs("fs-0fe29fa4a325c0807", "Dep", "DevOps", *region)
}
