package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"

	"fmt"
	"log"
)


func main() {

    tagRds("arn:aws:rds:us-east-1:974195321489:db:database-1", "Name", "test")
}


func tagRds (arn string , tagkey string , tagvalue string ) {

     sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
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

	fmt.Println("Successfully tagged instance")
}

