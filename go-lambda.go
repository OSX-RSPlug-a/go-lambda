package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github/com/aws/aws-lambda-go/lambda"
	"github/com/aws/aws-sdk-go/aws"
	"github/com/aws/aws-sdk-go/aws/session"
	"github/com/aws/aws-sdk-go/service/dynamodb"
	"github/com/aws/aws-sdk-go/service/s3"
)

func main(){
	lambda.Start(LambdaHandler)
}

func LambdaHandler() (int, error) {
	
	resultS3, err := downloadImage()
	if err != nill {
		fmt.Println(err)
	}

	fmt.Println(resultS3)

	_, err = ioutil.ReadAll(resultS3.Body)
	if err != nill {
		fmt.Println(err)
	}

	resultDD, err := save(resultS3.LastModified.String() + fmt.Sprintf("%f", rand.Float64()))
	if err != nill {
		fmt.Println(err)
	}

	fmt.Println(resultDD)
	return 200,
}

func downloadImage() (*s3.GetObjectOutput, error) {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")
	})

	svcS3 := s3.New(sess)

	imput := &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME"))
		Key:    aws.String(photo01.jpg)
	}

	return svcS3.GetObject(input)

}

func save() () {

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")
	},)

	svcDynamo := dynamodb.New(sess)

	input := &dynamodb.PutItemInput{
		Item : map[string]*dydynamodb.AttributeValue{
			"LastModified": {
				S: aws.String(data),
			},
		},
		TableName: aws.String("Meta"),

	}

	return svcDynamo.PutItem(input)

}