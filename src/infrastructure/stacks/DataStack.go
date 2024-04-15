package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DataStackProps struct {
	awscdk.StackProps
	// Add any custom fields if needed
}

func NewDataStack(scope constructs.Construct, id string, props *DataStackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, &props.StackProps)

	// Table for Weather API responses
	awsdynamodb.NewTable(stack, jsii.String("WeatherTable"), &awsdynamodb.TableProps{
		TableName:   jsii.String("WeatherData"),
		BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("requestId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})

	// Table for GitHub API responses
	awsdynamodb.NewTable(stack, jsii.String("GithubTable"), &awsdynamodb.TableProps{
		TableName:   jsii.String("GithubData"),
		BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("requestId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})

	// Table for Google Geocoding API responses
	awsdynamodb.NewTable(stack, jsii.String("GoogleGeoTable"), &awsdynamodb.TableProps{
		TableName:   jsii.String("GoogleGeocodingData"),
		BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("requestId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})

	// Table for Stripe transactions
	awsdynamodb.NewTable(stack, jsii.String("StripeTable"), &awsdynamodb.TableProps{
		TableName:   jsii.String("StripeTransactions"),
		BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("transactionId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})

	return stack
}
