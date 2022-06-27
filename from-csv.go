package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "flag"
	"pi-square-io/go-aws-tags/functions"
)

type empData struct {
    resource string
    id string
    tagkey string
    tagvalue string
}

var (
	region   = flag.String("region", "us-east-1", "your AWS region")
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
            id: line[1],
			tagkey: line[2],
			tagvalue: line[3],
        }
        fmt.Println(emp.resource + " " + emp.id + " " + emp.tagkey + " " + emp.tagvalue)

        if emp.resource == "ec2" {
            var instanceId = []string{emp.id}
		    functions.TagEc2(instanceId, emp.tagkey, emp.tagvalue, *region)
        }
    }
}