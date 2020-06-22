package main

import (
	"flag"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

	// Create S3 service client
	svc := s3.New(sess)

	result, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		log.Println("Failed to list buckets", err)
		return
	}

	log.Println("Buckets:")

	for _, bucket := range result.Buckets {
		log.Printf("%s : %s\n", *bucket.Name, bucket.CreationDate)
	}
}
