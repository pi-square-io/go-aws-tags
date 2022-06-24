package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"

	"fmt"
	"log"
)

func TagElb(elbArn []string, tagkey string, tagvalue string) {


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
				Key:   aws.String(tagkey),
				Value: aws.String(tagvalue),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for ELB", elbArns, errtag)
		return
	}

	fmt.Println("Successfully tagged instance")
}
