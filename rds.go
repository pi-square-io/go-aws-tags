package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"

	"fmt"
	"log"
)

func main() {

	region := "us-east-1"
	arn := "arn:aws:rds:us-east-1:974195321489:db:database-test-tags"

	tag_key := "Name"
	tag_value := "GoTest"

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	// Create RDS service client
	svc := rds.New(sess)

	// Add tags to the created instance
	_, errtag := svc.AddTagsToResource(&rds.AddTagsToResourceInput{
		ResourceName: aws.String(arn),
		Tags: []*rds.Tag{
			{
				Key:   aws.String(tag_key),
				Value: aws.String(tag_value),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for rds", arn, errtag)
		return
	}

	fmt.Println("Successfully tagged instance")
}
