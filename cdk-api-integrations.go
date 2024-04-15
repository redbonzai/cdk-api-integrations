package main

import (
	"cdk-api-integrations/src/infrastructure/stacks"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	app := awscdk.NewApp(nil)

	stacks.NewDataStack(app, "DataStack", &stacks.DataStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	stacks.NewLambdaStack(app, "LambdaStack", &stacks.LambdaStackProps{
		StackProps: awscdk.StackProps{
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
