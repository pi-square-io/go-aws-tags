#  GO - AWS Tag Resources

- these scripts allow adding tags to AWS resources from CSV file, from command line  using GoLang SDK.

# Technology

- [AWS](https://aws.amazon.com/what-is-aws/) Amazon Web Services (AWS) is the world’s most comprehensive and broadly adopted cloud platform, offering over 200 fully featured services from data centers globally. Millions of customers—including the fastest-growing startups, largest enterprises, and leading government agencies—are using AWS to lower costs, become more agile, and innovate faster.
- [GoLang](https://go.dev/doc/) The Go programming language is an open source project to make programmers more productive.


# Getting Started

Ensure that you have installed the following tools in your Mac or Linux or Windows Laptop before start working with this script and run go run.

1. [aws cli](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html)
2. [Go](https://go.dev/doc/install)

### Create AWS Profile

Add the block below to your .aws/credentials file and change the profile_name, aws_access_key_id and aws_secret_access_key with yours.

```shell script
[profile-name]
aws_access_key_id = XXXXXXX
aws_secret_access_key = XXXXXXXXXXXXXXXXX
```

### Clone the repo

```shell script
$ git clone git@github.com:pi-square-io/go-aws-tags.git
```

### Bulk tag resources

```shell script
$  go run bulk.go --resources resourceType-1 --ids id-or-arn-or-name-of-resource --tagKey key-of-tag --tagValue value-of-tag --resources resourceType-2 --ids id-or-arn-or-name-of-resource --tagKey key-of-tag --tagValue value-of-tag --resources resourceType-N --ids id-or-arn-or-name-of-resource --tagKey key-of-tag --tagValue value-of-tag
```

### Tag multiple resources from a CSV file

- add data in ./data.csv
```shell script
resourcetype-1,id-or-arn-or-name,tagkey,tagvalue
resourcetype-2,id-or-arn-or-name,tagkey,tagvalue
.
.
.
resourcetype-N,id-or-arn-or-name,tagkey,tagvalue
```

- Run go run 

```shell script
$ go run from-csv.go
```

### tag a single resource with a single tag

```shell script
$ go run main.go --resources resourceType --ids id-or-arn-or-name-of-resource --tagKey key-of-tag --tagValue value-of-tag
```
