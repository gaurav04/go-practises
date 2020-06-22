package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func DeleteAlarm(sess *session.Session, alarmName *string) error {

	svc := cloudwatch.New(sess)

	_, err := svc.DeleteAlarms(&cloudwatch.DeleteAlarmsInput{AlarmNames: []*string{alarmName}})

	if err != nil {
		return err
	}

	return nil
}

func main() {

	var region, accessKey, secretKey, alarmName string
	flag.StringVar(&accessKey, "accesskey", "", "access key")
	flag.StringVar(&secretKey, "secretkey", "", "secret-key")
	flag.StringVar(&region, "region", "us-west-2", "aws region")
	flag.StringVar(&alarmName, "alarmName", "", "alarm name to be deleted")

	flag.Parse()

	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region), Credentials: credentials.NewStaticCredentials(accessKey, secretKey, "")})

	err := DeleteAlarm(sess, &alarmName)
	if err != nil {
		fmt.Println("Could not delete alarm " + alarmName)
	} else {
		fmt.Println("Deleted alarm " + alarmName)
	}
}
