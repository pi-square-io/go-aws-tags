package functions

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"

	"fmt"
	"log"
)

func TagRds(arn string, tagkey string, tagvalue string, region string) {

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	svc := rds.New(sess)

	// Add tags to the created instance
	_, errtag := svc.AddTagsToResource(&rds.AddTagsToResourceInput{
		ResourceName: aws.String(arn),
		Tags: []*rds.Tag{
			{
				Key:   aws.String(tagkey),
				Value: aws.String(tagvalue),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for rds", arn, errtag)
		return
	}

	fmt.Println("Successfully tagged RDS")
}