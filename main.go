package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"

	"tarun-cdk-go/stacks"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	stacks.NewS3Stack(app, "S3Stack", nil)
	stacks.NewSQSStack(app, "SQSStack", nil)
	// stacks.NewEC2Stack(app, "EC2Stack", nil)

	app.Synth(nil)
}
