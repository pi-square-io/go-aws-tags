package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"fmt"
	"log"
)

func main() {

	region := "us-east-1"
	resources := []string{
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
	_, errtag := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: aws.StringSlice(resources),
		Tags: []*ec2.Tag{
			{
				Key:   aws.String(tag_key),
				Value: aws.String(tag_value),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for instance", "i-06114e47b6fc36fd2", errtag)
		return
	}

	fmt.Println("Successfully tagged instance")
}
