package functions

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/efs"

	"fmt"
	"log"
)

func TagEfs(fsId string, tagkey string, tagvalue string, region string) {

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	svc := efs.New(sess)

	// Add tags to the created instance
	_, errtag := svc.CreateTags(&efs.CreateTagsInput{
		FileSystemId: aws.String(fsId),
		Tags: []*efs.Tag{
			{
				Key:   aws.String(tagkey),
				Value: aws.String(tagvalue),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for EFS", errtag)
		return
	}

	fmt.Println("Successfully tagged EFS")
}