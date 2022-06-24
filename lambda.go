package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"

	"fmt"
	"log"
)

func main() {

	region := "us-east-1"
	resource := ""

	tag_key := "Name"
	tag_value := "GoTest"

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	// Create EC2 service client
	svc := lambda.New(sess)

	// Add tags to the created instance
	_, errtag := svc.TagResource(&lambda.TagResourceInput{
		Resource: aws.String(resource),
		Tags: []*lambda.Tag{
			{
				Key:   aws.String(tag_key),
				Value: aws.String(tag_value),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for instance", errtag)
		return
	}

	fmt.Println("Successfully tagged instance")
}
