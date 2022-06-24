package functions

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"fmt"
)

func TagS3(bucket string, tagName1 string, tagValue1 string, region string) {
	// Pre-defined values

	// from the shared credentials file. (~/.aws/credentials).
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Create S3 service client
	svc := s3.New(sess)

	// Create input for PutBucket method
	putInput := &s3.PutBucketTaggingInput{
		Bucket: aws.String(bucket),
		Tagging: &s3.Tagging{
			TagSet: []*s3.Tag{
				{
					Key:   aws.String(tagName1),
					Value: aws.String(tagValue1),
				},
				//                 {
				//                     Key:   aws.String(tagName2),
				//                     Value: aws.String(tagValue2),
				//                 },
			},
		},
	}

	_, err = svc.PutBucketTagging(putInput)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Now show the tags
	// Create input for GetBucket method
	getInput := &s3.GetBucketTaggingInput{
		Bucket: aws.String(bucket),
	}

	result, err := svc.GetBucketTagging(getInput)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	numTags := len(result.TagSet)

	if numTags > 0 {
		fmt.Println("Found", numTags, "Tag(s):")
		fmt.Println("")

		for _, t := range result.TagSet {
			fmt.Println("  Key:  ", *t.Key)
			fmt.Println("  Value:", *t.Value)
			fmt.Println("")
		}
	} else {
		fmt.Println("Did not find any tags")
	}
}
