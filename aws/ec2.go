package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

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
	// Create new EC2 client
	ec2Svc := ec2.New(sess)
	result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success", result)
	}

	for _, r := range result.Reservations {
		fmt.Println("Reservation ID: " + *r.ReservationId)
		fmt.Println("Instance IDs:")
		for _, i := range r.Instances {
			fmt.Println("   " + *i.InstanceId)
		}

		fmt.Println("")
	}

}
