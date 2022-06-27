package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

type empData struct {
    Name string
    Age string
    City string
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
            Name: line[0],
            Age: line[1],
			City: line[2],
        }
        fmt.Println(emp.Name + " " + emp.Age + " " + emp.City)
    }