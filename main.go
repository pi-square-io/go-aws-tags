package main

import (
	"flag"
	"pi-square-io/go-aws-tags/functions"
)

var (
	region = flag.String("region", "us-east-1", "your AWS region")
)

func main() {

	functions.TagRds("arn:aws:rds:us-east-1:974195321489:db:database-1", "Name", "test", *region)
//
	var instanceId = []string{"i-0ee77517f91d44f71"}
	functions.TagEc2(instanceId, "Name", "test", *region)

// 	tagS3("my-unique-435frzqrf", "Name", "test", *region)

    var albArns = []string{"i-0ee77517f91d44f71"}
	functions.TagElb(albArns, "Name", "test", *region)
}

//
//
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
