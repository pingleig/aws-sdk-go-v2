// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the refresh status of the AWS Trusted Advisor checks that have the
// specified check IDs. You can get the check IDs by calling the
// DescribeTrustedAdvisorChecks operation. Some checks are refreshed automatically,
// and you can't return their refresh statuses by using the
// DescribeTrustedAdvisorCheckRefreshStatuses operation. If you call this operation
// for these checks, you might see an InvalidParameterValue error.
//
// * You must have
// a Business or Enterprise Support plan to use the AWS Support API.
//
// * If you call
// the AWS Support API from an account that does not have a Business or Enterprise
// Support plan, the SubscriptionRequiredException error message appears. For
// information about changing your support plan, see AWS Support
// (http://aws.amazon.com/premiumsupport/).
func (c *Client) DescribeTrustedAdvisorCheckRefreshStatuses(ctx context.Context, params *DescribeTrustedAdvisorCheckRefreshStatusesInput, optFns ...func(*Options)) (*DescribeTrustedAdvisorCheckRefreshStatusesOutput, error) {
	if params == nil {
		params = &DescribeTrustedAdvisorCheckRefreshStatusesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTrustedAdvisorCheckRefreshStatuses", params, optFns, addOperationDescribeTrustedAdvisorCheckRefreshStatusesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTrustedAdvisorCheckRefreshStatusesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTrustedAdvisorCheckRefreshStatusesInput struct {

	// The IDs of the Trusted Advisor checks to get the status. If you specify the
	// check ID of a check that is automatically refreshed, you might see an
	// InvalidParameterValue error.
	//
	// This member is required.
	CheckIds []string
}

// The statuses of the Trusted Advisor checks returned by the
// DescribeTrustedAdvisorCheckRefreshStatuses operation.
type DescribeTrustedAdvisorCheckRefreshStatusesOutput struct {

	// The refresh status of the specified Trusted Advisor checks.
	//
	// This member is required.
	Statuses []types.TrustedAdvisorCheckRefreshStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTrustedAdvisorCheckRefreshStatusesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTrustedAdvisorCheckRefreshStatuses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTrustedAdvisorCheckRefreshStatuses{}, middleware.After)
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
	if err = addOpDescribeTrustedAdvisorCheckRefreshStatusesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrustedAdvisorCheckRefreshStatuses(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTrustedAdvisorCheckRefreshStatuses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "DescribeTrustedAdvisorCheckRefreshStatuses",
	}
}
