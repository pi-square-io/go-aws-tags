package functions

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"

	"fmt"
	"log"
)

func TagLambda(resourceArn, tagkey, tagvalue, region string) {

    sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	// Create EC2 service client
	svc := lambda.New(sess)

	// Add tags to the created instance
	_, errtag := svc.TagResource(&lambda.TagResourceInput{
		Resource: aws.String(resourceArn),
		Tags: map[string]*string{
				tagkey: aws.String(tagvalue),
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for lambda", errtag)
		return
	}

	fmt.Println("Successfully tagged lambda")
}
