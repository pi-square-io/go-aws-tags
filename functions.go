package main

import (
// 	"flag"
// 	"fmt"
// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/ec2"
// 	"github.com/aws/aws-sdk-go/service/rds"
// 	"github.com/aws/aws-sdk-go/service/s3"
// 	"log"
)

var (
	region = flag.String("region", "us-east-1", "your AWS region")
)

func main() {

// 	tagRds("arn:aws:rds:us-east-1:974195321489:db:database-1", "Name", "test")
//
// 	var instanceId = []string{"i-0ee77517f91d44f71"}
// 	tagEc2(instanceId, "Name", "test")

// 	tagS3("my-unique-435frzqrf", "Name", "test")

    var albArns = []string{"i-0ee77517f91d44f71"}
	TagElb(albArns, "Name", "test")
}

// func tagRds(arn string, tagkey string, tagvalue string) {
//
// 	sess, _ := session.NewSession(&aws.Config{
// 		Region: aws.String(*region)},
// 	)
// 	svc := rds.New(sess)
//
// 	// Add tags to the created instance
// 	_, errtag := svc.AddTagsToResource(&rds.AddTagsToResourceInput{
// 		ResourceName: aws.String(arn),
// 		Tags: []*rds.Tag{
// 			{
// 				Key:   aws.String(tagkey),
// 				Value: aws.String(tagvalue),
// 			},
// 		},
// 	})
// 	if errtag != nil {
// 		log.Println("Could not create tags for rds", arn, errtag)
// 		return
// 	}
//
// 	fmt.Println("Successfully tagged RDS")
// }
//
// func tagEc2(id []string, tagkey string, tagvalue string) {
//
// 	sess, _ := session.NewSession(&aws.Config{
// 		Region: aws.String(*region)},
// 	)
// 	svc := ec2.New(sess)
//
// 	// Add tags to the created instance
// 	_, errtag := svc.CreateTags(&ec2.CreateTagsInput{
// 		Resources: aws.StringSlice(id),
// 		Tags: []*ec2.Tag{
// 			{
// 				Key:   aws.String(tagkey),
// 				Value: aws.String(tagvalue),
// 			},
// 		},
// 	})
// 	if errtag != nil {
// 		log.Println("Could not create tags for instance", errtag)
// 		return
// 	}
//
// 	fmt.Println("Successfully tagged instance")
// }
//
// func tagS3(bucket string, tagName1 string, tagValue1 string) {
// 	// Pre-defined values
//
// 	// from the shared credentials file. (~/.aws/credentials).
// 	sess, err := session.NewSession(&aws.Config{
// 		Region: aws.String(*region)},
// 	)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
//
// 	// Create S3 service client
// 	svc := s3.New(sess)
//
// 	// Create input for PutBucket method
// 	putInput := &s3.PutBucketTaggingInput{
// 		Bucket: aws.String(bucket),
// 		Tagging: &s3.Tagging{
// 			TagSet: []*s3.Tag{
// 				{
// 					Key:   aws.String(tagName1),
// 					Value: aws.String(tagValue1),
// 				},
// 				//                 {
// 				//                     Key:   aws.String(tagName2),
// 				//                     Value: aws.String(tagValue2),
// 				//                 },
// 			},
// 		},
// 	}
//
// 	_, err = svc.PutBucketTagging(putInput)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
//
// 	// Now show the tags
// 	// Create input for GetBucket method
// 	getInput := &s3.GetBucketTaggingInput{
// 		Bucket: aws.String(bucket),
// 	}
//
// 	result, err := svc.GetBucketTagging(getInput)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
//
// 	numTags := len(result.TagSet)
//
// 	if numTags > 0 {
// 		fmt.Println("Found", numTags, "Tag(s):")
// 		fmt.Println("")
//
// 		for _, t := range result.TagSet {
// 			fmt.Println("  Key:  ", *t.Key)
// 			fmt.Println("  Value:", *t.Value)
// 			fmt.Println("")
// 		}
// 	} else {
// 		fmt.Println("Did not find any tags")
// 	}
// }
