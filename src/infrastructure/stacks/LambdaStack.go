package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type LambdaStackProps struct {
	awscdk.StackProps
	WeatherTable awsdynamodb.ITable // Assuming DynamoDB tables are needed
}

func NewLambdaStack(scope constructs.Construct, id string, props *LambdaStackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, &props.StackProps)

	// Define Lambda function for Weather API
	weatherLambda := awslambda.NewFunction(stack, jsii.String("WeatherLambda"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("weather.handler"),
		Code:    awslambda.Code_FromAsset(jsii.String("src/infrastructure/lambda"), nil),
	})

	// Define Lambda function for GitHub API
	githubLambda := awslambda.NewFunction(stack, jsii.String("GithubLambda"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("github.handler"),
		Code:    awslambda.Code_FromAsset(jsii.String("src/infrastructure/lambda"), nil),
	})

	// Define Lambda function for Google Geocoding API
	googleLambda := awslambda.NewFunction(stack, jsii.String("GoogleGeoLambda"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("googlegeo.handler"),
		Code:    awslambda.Code_FromAsset(jsii.String("src/infrastructure/lambda"), nil),
	})

	// Define Lambda function for Stripe API
	stripeLambda := awslambda.NewFunction(stack, jsii.String("StripeLambda"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("stripe.handler"),
		Code:    awslambda.Code_FromAsset(jsii.String("src/infrastructure/lambda"), nil),
	})

	// Create the API Gateway
	api := awsapigateway.NewRestApi(stack, jsii.String("ApiGateway"), &awsapigateway.RestApiProps{
		RestApiName: jsii.String("LambdaIntegrationApi"),
	})

	// Weather API integration
	weatherResource := api.Root().AddResource(jsii.String("weather"), nil)
	weatherResource.AddMethod(jsii.String("GET"), awsapigateway.NewLambdaIntegration(weatherLambda, nil), nil)

	// GitHub API integration
	githubResource := api.Root().AddResource(jsii.String("github"), nil)
	githubResource.AddMethod(jsii.String("GET"), awsapigateway.NewLambdaIntegration(githubLambda, nil), nil)

	// Google Geocoding API integration
	googleGeoResource := api.Root().AddResource(jsii.String("geocoding"), nil)
	googleGeoResource.AddMethod(jsii.String("GET"), awsapigateway.NewLambdaIntegration(googleLambda, nil), nil)

	// Stripe API integration
	stripeResource := api.Root().AddResource(jsii.String("stripe"), nil)
	stripeResource.AddMethod(jsii.String("POST"), awsapigateway.NewLambdaIntegration(stripeLambda, nil), nil)

	return stack
}
