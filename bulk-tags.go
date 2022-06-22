package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"

	"fmt"
	"log"
)

func main() {

	region := "us-east-1"
	instances := []string{
		"i-06114e47b6fc36fd2",
	}

	tag_key := "Name"
	tag_value := "GoTest"

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	// Create EC2 service client
	svc := resourcegroupstaggingapi.New(sess)

	// Add tags to the created instance
	_, errtag := svc.TagResources(&resourcegroupstaggingapi.TagResourcesInput{
		ResourceARNList: aws.StringSlice(instances),
		Tags: map[string]string{
			{
// 				Key:   aws.String(tag_key),
// 				Value: aws.String(tag_value),
                "key": "Name"
                value
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for instance", instances, errtag)
		return
	}

	fmt.Println("Successfully tagged instance")
}


####
for i in
    service= a[0]
    if sERVICE == ec2:

ec2 id key tag_valus
ec2 id key tag_valus
ec2b
s3 id
