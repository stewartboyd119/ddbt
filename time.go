package ddbt

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pkg/errors"
)

type TimeRFC3339 struct {
	T time.Time
}

func NewTimeRfc3339(t time.Time) TimeRFC3339 {
	return TimeRFC3339{
		T: t,
	}
}

func (i TimeRFC3339) MarshalDynamoDBAttributeValue(t time.Time, format string, av *dynamodb.AttributeValue) error {
	return marshalDynamoDBAttributeValue(i.T, time.RFC3339Nano, av)
}

func (i *TimeRFC3339) UnmarshalDynamoDBAttributeValue(av *dynamodb.AttributeValue) error {
	t, err := unmarshalDynamoDBAttributeValue(time.RFC3339Nano, av)
	i.T = t
	return err

}

func marshalDynamoDBAttributeValue(t time.Time, format string, av *dynamodb.AttributeValue) error {
	if av == nil {
		return errors.New("nil attributeValue")
	}
	av.S = aws.String(t.Format(format))
	return nil
}

func unmarshalDynamoDBAttributeValue(format string, av *dynamodb.AttributeValue) (time.Time, error) {
	if av == nil {
		return time.Time{}, errors.New("nil attributeValue")
	}
	t, err := time.Parse(format, aws.StringValue(av.S))
	if err != nil {
		return time.Time{}, errors.Wrap(err, "error parsing time arrtibuteValue")
	}
	return t, nil

}
