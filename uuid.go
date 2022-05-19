package ddbt

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type UUID struct {
	U uuid.UUID
}

func New(u uuid.UUID) UUID {
	return UUID{
		U: u,
	}
}

func (i UUID) MarshalDynamoDBAttributeValue(av *dynamodb.AttributeValue) error {
	if av == nil {
		return errors.New("nil attribute")
	}
	av.S = aws.String(i.U.String())
	return nil
}

func (i *UUID) UnmarshalDynamoDBAttributeValue(av *dynamodb.AttributeValue) error {
	if av == nil {
		return errors.New("nil attributeValue")
	}
	t := aws.StringValue(av.S)

	u, err := uuid.Parse(t)

	if err != nil {
		return errors.Wrap(err, "error parsing uuid attribute value")
	}
	i.U = u
	return nil

}
