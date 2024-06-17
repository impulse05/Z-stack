package repository

import (
	"OrderManagementService/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type OrderRepository interface {
	SaveOrder(order *model.Order) error
	GetOrder(orderID string) (*model.Order, error)
	UpdateOrder(order *model.Order) error
	GetAllOrders() ([]*model.Order, error)
}

type orderRepository struct {
	db *dynamodb.DynamoDB
}

func NewOrderRepository() OrderRepository {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Endpoint:    aws.String("http://localhost:8000"),
		Credentials: credentials.NewStaticCredentials("local", "local", ""),
	}))
	return &orderRepository{
		db: dynamodb.New(sess),
	}
}

func (r *orderRepository) SaveOrder(order *model.Order) error {
	av, err := dynamodbattribute.MarshalMap(order)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Orders"),
	}
	_, err = r.db.PutItem(input)
	return err
}

func (r *orderRepository) GetOrder(orderID string) (*model.Order, error) {
	result, err := r.db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Orders"),
		Key: map[string]*dynamodb.AttributeValue{
			"OrderID": {
				S: aws.String(orderID),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}
	order := new(model.Order)
	err = dynamodbattribute.UnmarshalMap(result.Item, order)
	return order, err
}

func (r *orderRepository) UpdateOrder(order *model.Order) error {
	av, err := dynamodbattribute.MarshalMap(order)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Orders"),
	}
	_, err = r.db.PutItem(input)
	return err
}

func (r *orderRepository) GetAllOrders() ([]*model.Order, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Orders"),
	}
	result, err := r.db.Scan(input)
	if err != nil {
		return nil, err
	}
	orders := []*model.Order{}
	for _, item := range result.Items {
		order := new(model.Order)
		err = dynamodbattribute.UnmarshalMap(item, order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
