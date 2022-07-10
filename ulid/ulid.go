package ddbt

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/oklog/ulid/v2"
	"github.com/pkg/errors"
)

type ULID struct {
	U ulid.ULID
}

func New(u ulid.ULID) ULID {
	return ULID{
		U: u,
	}
}

func (i ULID) MarshalDynamoDBAttributeValue(av *dynamodb.AttributeValue) error {
	if av == nil {
		return errors.New("nil attribute")
	}
	av.S = aws.String(i.U.String())
	return nil
}

func (i *ULID) UnmarshalDynamoDBAttributeValue(av *dynamodb.AttributeValue) error {
	if av == nil {
		return errors.New("nil attributeValue")
	}
	t := aws.StringValue(av.S)

	u, err := ulid.Parse(t)

	if err != nil {
		return errors.Wrap(err, "error parsing ulid attribute value")
	}
	i.U = u
	return nil

}

func (i ULID) String() string {
	return i.U.String()
}
