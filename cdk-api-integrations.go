package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DataStackProps struct {
	awscdk.StackProps
}

type LambdaStackProps struct {
	awscdk.StackProps
}

func NewDataStack(scope constructs.Construct, id string, props *DataStackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, &props.StackProps)
	// DynamoDB tables will be defined here.
	return stack
}

func NewLambdaStack(scope constructs.Construct, id string, props *LambdaStackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, &props.StackProps)
	// Lambda functions will be defined here.
	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewDataStack(app, "DataStack", &DataStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	NewLambdaStack(app, "LambdaStack", &LambdaStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String("your-aws-account-id"),
		Region:  jsii.String("your-aws-region"),
	}
}
