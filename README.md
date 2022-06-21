#  GO - AWS Tag Resources

- these scripts allow getting tags of resources, cost and usage of a resource, dimension values, usage forecast and untag a resource from AWS Cost Explorer using GoLang SDK.

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

### list tags of resources

- Set start and end date and the region

```shell script
func main() {

#       //Must be in YYYY-MM-DD Format
        start := "2022-05-01"
        end := "2022-05-27"

        region := "us-east-1"

#       // Initialize a session
#       sess, err := session.NewSession(&aws.Config{
#               Region: aws.String(region)},
        )
```
- Run go run to get tags of resources

```shell script
$ go run tags.go
```

### untag a resource

- Set resourceArn, list of key tags and the region

```shell script
func main() {

        resource := "arn:aws:resource:region:acount_id:resource_type/resource_name"
        tags := []string{
                "tag-key1",
                "tag-key2",
        }
        region := "us-east-1"

#       // Initialize a session
#       sess, err := session.NewSession(&aws.Config{
#               Region: aws.String(region)},
        )

```
- Run go run to untag a resource

```shell script
$ go run untag.go

```
### Get cost and usage of a resource

- Set start and end date, granularity, metrics and the region

```shell script
func main() {

        //Must be in YYYY-MM-DD Format
        start := "2022-01-01"
        end := "2022-05-01"

        granularity := "MONTHLY"
        metrics := []string{
                "BlendedCost",
                "UnblendedCost",
                "UsageQuantity",
        }

        region := "us-east-1"

#       // Initialize a session
#       sess, err := session.NewSession(&aws.Config{
#               Region: aws.String(region)},
        )
```

- Run go run get cost and usage of a resource

```shell script
$ go run cost-usage.go
```


### get dimension values

- Set start and end date, dimension, and the region

```shell script
func main() {

        //Must be in YYYY-MM-DD Format
        start := "2022-01-01"
        end := "2022-05-01"

    dimension := "USAGE_TYPE"

        region := "us-east-1"

#       // Initialize a session
#       sess, err := session.NewSession(&aws.Config{
#               Region: aws.String(region)},
        )
```

- Run go run get dimension values

```shell script
$ go run dimension.go
```

### get  get usage forecast

- Set start and end date, granularity, metrics, and the region

```shell script
func main() {

        //Must be in YYYY-MM-DD Format
        start := "2022-05-27" // start date must be in present or future
        end := "2022-06-30"   // end date must be in future

        granularity := "MONTHLY"
        metric := "UNBLENDED_COST"

        region := "us-east-1"

#       // Initialize a session
#       sess, err := session.NewSession(&aws.Config{
#               Region: aws.String(region)},
        )
```

- Run go run get usage forecast

```shell script
$ go run forecast.go
