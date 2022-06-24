package functions

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"fmt"
	"log"
)

func TagEc2(id []string, tagkey string, tagvalue string, region string) {

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
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
