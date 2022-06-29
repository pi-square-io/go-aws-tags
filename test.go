package main

import "flag"
// import "fmt"
import "pi-square-io/go-aws-tags/functions"

type arrayFlags []string

func (i *arrayFlags) String() string {
    return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
    *i = append(*i, value)
    return nil
}

var resources arrayFlags
var ids arrayFlags
var tagKey arrayFlags
var tagValue arrayFlags
var region   = flag.String("region", "us-east-1", "your AWS region")

func main() {
    flag.Var(&resources, "resources", "Some description for this param.")
    flag.Var(&ids, "ids", "Some description for this param.")
    flag.Var(&tagKey, "tagKey", "Some description for this param.")
    flag.Var(&tagValue, "tagValue", "Some description for this param.")
    flag.Parse()
//     fmt.Printf("Resrouces %s\n", resources)
//     fmt.Printf("ids %s\n", ids)
//     fmt.Printf("tagKey %s\n", tagKey)
//     fmt.Printf("tagValue %s\n", tagValue)
//     fmt.Printf("size %d\n", len(resources))
//     fmt.Printf("first resource %s\n",resources[0])

    for i:=0;i<len(resources);i++ {
        if resources[i] == "ec2" {
            var instanceId = []string{ids[i]}
		    functions.TagEc2(instanceId, tagKey[i], tagValue[i], *region)
         }

          if resources[i] == "rds" {
		    functions.TagRds(ids[i], tagKey[i], tagValue[i], *region)
         }

          if resources[i] == "s3" {
		    functions.TagS3(ids[i], tagKey[i], tagValue[i], *region)
         }

          if resources[i] == "lambda" {
		    functions.TagLambda(ids[i], tagKey[i], tagValue[i], *region)
         }
          if resources[i] == "efs" {
		    functions.TagEfs(ids[i], tagKey[i], tagValue[i], *region)
         }
          if resources[i] == "rds" {
		    functions.TagRds(ids[i], tagKey[i], tagValue[i], *region)
         }
        }
        
    }



