# CDK API INTEGRATIONS
> This project is a proof of concept using CDK to build a scalable API that integrates 
> with multiple APIs and SDKs (SDKs forthcoming).

## Description
This project automates the deployment of serverless applications using AWS services including Lambda, DynamoDB, and API Gateway. It integrates with external APIs such as OpenWeatherMap, GitHub, Google Geocoding, and Stripe to demonstrate a multifaceted serverless architecture.

## SETUP
Obtain API keys from:
- **Google Geocoder service**
- **OpenWeatherMap API**
- **GitHub API** (this also requires a personal access token)
- **Stripe API** (they will also give you a private key to use)

### CONVERT THE ENV.example file into .env
```terminal
cp env.example .env
```
Enter those values into the .env file. Then run the following command:
```bash
cdk bootstrap && cdk synth && cdk deploy --all
```
We use the `--all` command to deploy all the stacks in the project. Otherwise, you'd enter only the stack that you want to deploy.

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Building and Running the Project

### Building the Project
Compile the Go application to ensure all components are correctly set up:
```bash
go build -o myapplication
```

### Running Locally
For local testing and debugging:
```bash
./myapplication
```

### Deploying to AWS
Deploy your application to the AWS Cloud using the CDK toolkit:
```bash
cdk deploy --all
```

## Useful Commands

* `cdk deploy`      - Deploy this stack to your default AWS account/region.
* `cdk diff`        - Compare deployed stack with current state.
* `cdk synth`       - Emits the synthesized CloudFormation template.
* `go test`         - Run unit tests.

## Additional Resources
- [AWS CDK Documentation](https://docs.aws.amazon.com/cdk/latest/guide/home.html)
- [Troubleshooting guide]

## Contact Information
For help and support, reach out to `support@yourdomain.com`.

---

### We've Provided a Makefile

Including a Makefile can streamline the build and deployment process, making it easier to manage complex commands. Here's a simple Makefile template:

```makefile
.PHONY: install build deploy

install:
	npm install -g aws-cdk
	go mod tidy

build:
	go build -o cdk-api-integrations

deploy:
	cdk deploy --all

run:
	go run ./cdk-api-integrations.go
```

**Usage Instructions for Makefile:**
- **Install Dependencies**: `make install`
- **Build the Application**: `make build`
- **Deploy to AWS**: `make deploy`
- **Run Locally**: `make run`

This comprehensive README should guide new users through setting up and using your project effectively, while also ensuring that experienced developers have quick access to the necessary commands and configurations.
