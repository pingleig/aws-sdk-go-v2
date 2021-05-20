// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiate a manual deployment of the latest commit in a source code repository or
// the latest image in a source image repository to an AWS App Runner service. For
// a source code repository, App Runner retrieves the commit and builds a Docker
// image. For a source image repository, App Runner retrieves the latest Docker
// image. In both cases, App Runner then deploys the new image to your service and
// starts a new container instance. This is an asynchronous operation. On a
// successful call, you can use the returned OperationId and the ListOperations
// call to track the operation's progress.
func (c *Client) StartDeployment(ctx context.Context, params *StartDeploymentInput, optFns ...func(*Options)) (*StartDeploymentOutput, error) {
	if params == nil {
		params = &StartDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDeployment", params, optFns, addOperationStartDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDeploymentInput struct {

	// The Amazon Resource Name (ARN) of the App Runner service that you want to
	// manually deploy to.
	//
	// This member is required.
	ServiceArn *string
}

type StartDeploymentOutput struct {

	// The unique ID of the asynchronous operation that this request started. You can
	// use it combined with the ListOperations call to track the operation's progress.
	//
	// This member is required.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpStartDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpStartDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpStartDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDeployment(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opStartDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "StartDeployment",
	}
}
