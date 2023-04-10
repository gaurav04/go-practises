package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

// define main function
func main() {

	var region string
	var accessKey string
	var secretKey string
	flag.StringVar(&accessKey, "accesskey", "", "access key")
	flag.StringVar(&secretKey, "secretkey", "", "secret-key")
	flag.StringVar(&region, "region", "us-west-2", "aws region")
	//flag.parse()
	flag.Parse()

	sess, err := session.NewSession(&aws.Config{Region: aws.String(region), Credentials: credentials.NewStaticCredentials(accessKey, secretKey, "")})

	// Create RDS service client
	svc := rds.New(sess)

	result, err := svc.DescribeDBInstances(nil)
	if err != nil {
		fmt.Println(err.Error)
	}

	fmt.Println("Instances:")

	for _, d := range result.DBInstances {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(d.DBInstanceIdentifier), aws.TimeValue(d.InstanceCreateTime))
	}
}
