package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

type empData struct {
    resource string
    arn string
    tagkey string
    tagvalue string
}

func ec2 (arn, tagkey, tagvalue string ) {

    sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

    svc := ec2.New(sess)

	// Add tags to the created instance
	_, errtag := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: aws.StringSlice(resources),
		Tags: []*ec2.Tag{
			{
				Key:   aws.String(tagkey),
				Value: aws.String(tagvalue),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for instance", "i-06114e47b6fc36fd2", errtag)
		return
	}

	fmt.Println("Successfully tagged instance")
}

func main() {

    csvFile, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

    csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
        fmt.Println(err)
    }
    for _, line := range csvLines {
        emp := empData{
            resource: line[0],
            arn: line[1],
			tagkey: line[2],
			tagvalue: line[3],
        }
        fmt.Println(emp.resource + " " + emp.arn + " " + emp.tagkey + " " + emp.tagvalue)

        if emp.resource == ec2 {
        ret = ec2("sdfghjklkjhgf", "Name", "test")
        }

    }
}