package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"pi-square-io/go-aws-tags/functions"
)

type empData struct {
	resource string
	id       string
	tagkey   string
	tagvalue string
}

var (
	region = flag.String("region", "us-east-1", "your AWS region")
)

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
			id:       line[1],
			tagkey:   line[2],
			tagvalue: line[3],
		}

		if emp.resource == "ec2" {
			var instanceId = []string{emp.id}
			functions.TagEc2(instanceId, emp.tagkey, emp.tagvalue, *region)
		}

		if emp.resource == "rds" {
			functions.TagRds(emp.id, emp.tagkey, emp.tagvalue, *region)
		}

		if emp.resource == "s3" {
			functions.TagS3(emp.id, emp.tagkey, emp.tagvalue, *region)
		}

		if emp.resource == "alb" {
			var albArns = []string{emp.id}
			functions.TagElb(albArns, emp.tagkey, emp.tagvalue, *region)
		}

		if emp.resource == "lambda" {
			functions.TagLambda(emp.id, emp.tagkey, emp.tagvalue, *region)
		}

		if emp.resource == "efs" {
			functions.TagEfs(emp.id, emp.tagkey, emp.tagvalue, *region)
		}
	}
}
