// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops creating a solution version that is in a state of CREATE_PENDING or CREATE
// IN_PROGRESS. Depending on the current state of the solution version, the
// solution version state changes as follows:
//
// * CREATE_PENDING > CREATE_STOPPED
// or
//
// * CREATE_IN_PROGRESS > CREATE_STOPPING > CREATE_STOPPED
//
// You are billed for
// all of the training completed up until you stop the solution version creation.
// You cannot resume creating a solution version once it has been stopped.
func (c *Client) StopSolutionVersionCreation(ctx context.Context, params *StopSolutionVersionCreationInput, optFns ...func(*Options)) (*StopSolutionVersionCreationOutput, error) {
	if params == nil {
		params = &StopSolutionVersionCreationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopSolutionVersionCreation", params, optFns, addOperationStopSolutionVersionCreationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopSolutionVersionCreationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopSolutionVersionCreationInput struct {

	// The Amazon Resource Name (ARN) of the solution version you want to stop
	// creating.
	//
	// This member is required.
	SolutionVersionArn *string
}

type StopSolutionVersionCreationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopSolutionVersionCreationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopSolutionVersionCreation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopSolutionVersionCreation{}, middleware.After)
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
	if err = addOpStopSolutionVersionCreationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopSolutionVersionCreation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopSolutionVersionCreation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "StopSolutionVersionCreation",
	}
}
