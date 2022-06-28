package main

import (
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"fmt"
	"log"
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

	if *resource == "ec2" {
		var instanceId = []string{*id}
		TagEc2(instanceId, *tagkey, *tagvalue, *region)
	}
}

func TagEc2(id []string, tagkey string, tagvalue string, region string) {

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	svc := ec2.New(sess)

	// Add tags to the created instance
	_, errtag := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: aws.StringSlice(id),
		Tags: []*ec2.Tag{
			{
				Key:   aws.String(tagkey),
				Value: aws.String(tagvalue),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for instance", errtag)
		return
	}

	fmt.Println("Successfully tagged instance")
}
