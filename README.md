# CDK API INTEGRATIONS

This project is a proof of concept using CDK to build a scalable API that integrates with multiple
APIs and SDKs ( SDKs forthcoming).  

## SETUP
obtain api keys from: 
- Google Geocoder service
- Open weather map api
- Github API ( this also requires a personal access token )
- Stripe API (they will also give you a private key to use)

## CONVERT THE ENV.example file into .env
```terminal
cp env.example .env
```
Then, Enter those values into the .env file. 
Then run the following command : 
```bash
cdk bootstrap && cdk synth && cdk deploy --all
```
We use the `--all` command to deploy all the stacks in the project. Otherwise, you'd enter only the stack
that you want to deploy.


The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests
