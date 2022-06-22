package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
// 	"github.com/aws/aws-sdk-go/service/lambda"

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
	svc := ec2.New(sess)

	// Add tags to the created instance
	_, errtag := svc.CreateTags(&lambda.CreateTagsInput{
		Resources: aws.StringSlice(instances),
		Tags: []*ec2.Tag{
			{
				Key:   aws.String(tag_key),
				Value: aws.String(tag_value),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for instance", instances, errtag)
		return
	}

	fmt.Println("Successfully tagged instance")
}