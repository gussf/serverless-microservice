package db

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gussf/serverless-microservice/domain/interfaces"
)

type DynamoRepository struct {
	dynamoService *dynamodb.DynamoDB
}

func createDynamoSession() *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	sess.Config.Region = aws.String("us-east-1")
	return sess
}

func NewDynamoRepository() DynamoRepository {
	sess := createDynamoSession()
	_, err := sess.Config.Credentials.Get()
	if err != nil {
		fmt.Println(err.Error())
	}
	svc := dynamodb.New(sess)

	return DynamoRepository{dynamoService: svc}
}

func (d DynamoRepository) ListCars() ([]interfaces.CarDAO, error) {
	tableName := "Cars"

	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := d.dynamoService.Scan(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
		fmt.Println(err)
	}

	var cars []interfaces.CarDAO
	for _, i := range result.Items {
		car := interfaces.CarDAO{}

		err = dynamodbattribute.UnmarshalMap(i, &car)

		if err != nil {
			log.Fatalf("Got error unmarshalling: %s", err)
			fmt.Println(err)
		}

		cars = append(cars, car)
		fmt.Println(car)
	}

	return cars, nil
}

func (d DynamoRepository) InsertCar(car interfaces.CarDAO) error {
	return nil
}
