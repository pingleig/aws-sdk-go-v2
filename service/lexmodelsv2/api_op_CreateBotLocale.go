// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a locale in the bot. The locale contains the intents and slot types that
// the bot uses in conversations with users in the specified language and locale.
// You must add a locale to a bot before you can add intents and slot types to the
// bot.
func (c *Client) CreateBotLocale(ctx context.Context, params *CreateBotLocaleInput, optFns ...func(*Options)) (*CreateBotLocaleOutput, error) {
	if params == nil {
		params = &CreateBotLocaleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBotLocale", params, optFns, addOperationCreateBotLocaleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBotLocaleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBotLocaleInput struct {

	// The identifier of the bot to create the locale for.
	//
	// This member is required.
	BotId *string

	// The version of the bot to create the locale for. This can only be the draft
	// version of the bot.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the language and locale that the bot will be used in. The
	// string must match one of the supported locales. All of the intents, slot types,
	// and slots used in the bot must have the same locale. For more information, see
	// Supported languages
	// (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html).
	//
	// This member is required.
	LocaleId *string

	// Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent,
	// AMAZON.KendraSearchIntent, or both when returning alternative intents.
	// AMAZON.FallbackIntent and AMAZON.KendraSearchIntent are only inserted if they
	// are configured for the bot. For example, suppose a bot is configured with the
	// confidence threshold of 0.80 and the AMAZON.FallbackIntent. Amazon Lex returns
	// three alternative intents with the following confidence scores: IntentA (0.70),
	// IntentB (0.60), IntentC (0.50). The response from the PostText operation would
	// be:
	//
	// * AMAZON.FallbackIntent
	//
	// * IntentA
	//
	// * IntentB
	//
	// * IntentC
	//
	// This member is required.
	NluIntentConfidenceThreshold *float64

	// A description of the bot locale. Use this to help identify the bot locale in
	// lists.
	Description *string

	// The Amazon Polly voice ID that Amazon Lex uses for voice interaction with the
	// user.
	VoiceSettings *types.VoiceSettings
}

type CreateBotLocaleOutput struct {

	// The specified bot identifier.
	BotId *string

	// The status of the bot. When the status is Creating the bot locale is being
	// configured. When the status is Building Amazon Lex is building the bot for
	// testing and use. If the status of the bot is ReadyExpressTesting, you can test
	// the bot using the exact utterances specified in the bots' intents. When the bot
	// is ready for full testing or to run, the status is Built. If there was a problem
	// with building the bot, the status is Failed. If the bot was saved but not built,
	// the status is NotBuilt.
	BotLocaleStatus types.BotLocaleStatus

	// The specified bot version.
	BotVersion *string

	// A timestamp specifying the date and time that the bot locale was created.
	CreationDateTime *time.Time

	// The specified description of the bot locale.
	Description *string

	// The specified locale identifier.
	LocaleId *string

	// The specified locale name.
	LocaleName *string

	// The specified confidence threshold for inserting the AMAZON.FallbackIntent and
	// AMAZON.KendraSearchIntent intents.
	NluIntentConfidenceThreshold *float64

	// The Amazon Polly voice ID that Amazon Lex uses for voice interaction with the
	// user.
	VoiceSettings *types.VoiceSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateBotLocaleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBotLocale{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBotLocale{}, middleware.After)
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
	if err = addOpCreateBotLocaleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBotLocale(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBotLocale(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "CreateBotLocale",
	}
}
