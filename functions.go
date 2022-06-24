package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/ec2"
    "flag"
	"fmt"
	"log"
)

var (
	region = flag.String("region", "us-east-1", "your AWS region")
)

func main() {

    tagRds("arn:aws:rds:us-east-1:974195321489:db:database-1", "Name", "test")

    var instanceId = []string {"i-0ee77517f91d44f71",}
    tagEc2(instanceId, "Name", "test")
}


func tagRds (arn string , tagkey string , tagvalue string ) {

     sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(*region)},
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


func tagEc2 (id []string , tagkey string , tagvalue string ) {

     sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(*region)},
	)
   	svc := ec2.New(sess)

	// Add tags to the created instance
	_, errtag := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: aws.StringSlice(id),
		Tags: []*ec2.Tag{
			{
				Key:   aws.String(tagkey),
				Value: aws.String(tagvalue),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for instance", errtag)
		return
	}

	fmt.Println("Successfully tagged instance")
}
