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
    }
}