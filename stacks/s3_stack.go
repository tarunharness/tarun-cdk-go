package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewS3Stack(scope constructs.Construct, id string, props *awscdk.StackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, props)

	bucket := awss3.NewBucket(stack, jsii.String("sunny-cdk-test-bucket-03"), &awss3.BucketProps{
		BucketName:        jsii.String("sunny-cdk-test-bucket-03"),
		Versioned:         jsii.Bool(true),
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
		AutoDeleteObjects: jsii.Bool(true),
		Encryption:        awss3.BucketEncryption_S3_MANAGED,
	})

	awscdk.NewCfnOutput(stack, jsii.String("S3BucketARN"), &awscdk.CfnOutputProps{
		Value:       bucket.BucketArn(),
		Description: jsii.String("The ARN of the S3 bucket"),
		ExportName:  jsii.String("MyS3BucketARN"),
	})

	return stack
}
