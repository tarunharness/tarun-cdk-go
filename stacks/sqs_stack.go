package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewSQSStack(scope constructs.Construct, id string, props *awscdk.StackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, props)

	queueName := stack.Node().TryGetContext(jsii.String("queueName")).(string)
	retentionDays := stack.Node().TryGetContext(jsii.String("queueRetentionDays")).(float64)

	queue := awssqs.NewQueue(stack, jsii.String("MyQueue3"), &awssqs.QueueProps{
		QueueName:         jsii.String(queueName),
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(500)),
		RetentionPeriod:   awscdk.Duration_Days(jsii.Number(retentionDays)),
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
	})

	awscdk.NewCfnOutput(stack, jsii.String("SQSQueueARN"), &awscdk.CfnOutputProps{
		Value:       queue.QueueArn(),
		Description: jsii.String("The ARN of the SQS queue"),
		ExportName:  jsii.String("MySQSQueueARN"),
	})

	return stack
}
