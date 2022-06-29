package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
)

var (
	region    = flag.String("region", "us-east-1", "your AWS region")
	resource  = flag.String("resource", "ec2", "your resource")
	id        = flag.String("id", "id", "id,arn or name of resource")
	tagkey1   = flag.String("tagkey1", "tagkey1", "key of tag resource")
	tagvalue1 = flag.String("tagvalue1", "tagvalue1", "value of tag resource")
	tagkey2   = flag.String("tagkey2", "tagkey2", "key of tag resource")
	tagvalue2 = flag.String("tagvalue2", "tagvalue2", "value of tag resource")
	tagkey3   = flag.String("tagkey3", "tagkey3", "key of tag resource")
	tagvalue3 = flag.String("tagvalue3", "tagvalue3", "value of tag resource")
)

func main() {

	flag.Parse()

	if *resource == "ec2" {
		var instanceId = []string{*id}
		TagEc2(instanceId, *tagkey1, *tagvalue1, *tagkey2, *tagvalue2, *tagkey3, *tagvalue3, *region)
	}
}

func TagEc2(id []string, tagkey1 string, tagvalue1 string, tagkey2 string, tagvalue2 string, tagkey3 string, tagvalue3 string,region string) {

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	svc := ec2.New(sess)

	// Add tags to the created instance
	_, errtag := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: aws.StringSlice(id),
		Tags: []*ec2.Tag{
			{
				Key:   aws.String(tagkey1),
				Value: aws.String(tagvalue1),
			},
			{
				Key:   aws.String(tagkey2),
				Value: aws.String(tagvalue2),
			},
			{
				Key:   aws.String(tagkey3),
				Value: aws.String(tagvalue3),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for instance", errtag)
		return
	}

	fmt.Println("Successfully tagged instance")
}
