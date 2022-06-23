package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"

	"fmt"
	"log"
)

func main() {

	region := "us-east-1"
	elbArns := []string{
		"arn:aws:elasticloadbalancing:us-east-1:974195321489:loadbalancer/app/alb-test-tag/0c69155f72f7a598",
	}

	tag_key := "Name"
	tag_value := "GoTest"

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	// Create EC2 service client
	svc := elbv2.New(sess)

	// Add tags to the created instance
	_, errtag := svc.AddTags(&elbv2.AddTagsInput{
		ResourceArns: aws.StringSlice(elbArns),
		Tags: []*elbv2.Tag{
			{
				Key:   aws.String(tag_key),
				Value: aws.String(tag_value),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for ELB", elbArns ,errtag)
		return
	}

	fmt.Println("Successfully tagged instance")
}
