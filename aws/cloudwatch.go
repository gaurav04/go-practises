package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func ListAlarms(sess *session.Session) (*cloudwatch.DescribeAlarmsOutput, error) {

	svc := cloudwatch.New(sess)

	resp, err := svc.DescribeAlarms(nil)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func main() {

	var region, accessKey, secretKey string
	flag.StringVar(&accessKey, "accesskey", "", "access key")
	flag.StringVar(&secretKey, "secretkey", "", "secret-key")
	flag.StringVar(&region, "region", "us-west-2", "aws region")

	flag.Parse()

	sess, err := session.NewSession(&aws.Config{Region: aws.String(region), Credentials: credentials.NewStaticCredentials(accessKey, secretKey, "")})

	resp, err := ListAlarms(sess)
	if err != nil {
		fmt.Println("Got an error listing alarms:")
		fmt.Println(err)
		return
	}

	fmt.Println("Alarms:")
	for _, alarm := range resp.MetricAlarms {
		fmt.Println("    " + *alarm.AlarmName)
	}
}
