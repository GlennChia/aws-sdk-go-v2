// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a channel flow, a container for processors. Processors are AWS Lambda
// functions that perform actions on chat messages, such as stripping out
// profanity. You can associate channel flows with channels, and the processors in
// the channel flow then take action on all messages sent to that channel. This is
// a developer API. Channel flows process the following items:
//
// * New and updated
// messages
//
// * Persistent and non-persistent messages
//
// * The Standard message
// type
//
// Channel flows don't process Control or System messages. For more
// information about the message types provided by Chime SDK Messaging, refer to
// Message types
// (https://docs.aws.amazon.com/chime/latest/dg/using-the-messaging-sdk.html#msg-types)
// in the Amazon Chime developer guide.
func (c *Client) CreateChannelFlow(ctx context.Context, params *CreateChannelFlowInput, optFns ...func(*Options)) (*CreateChannelFlowOutput, error) {
	if params == nil {
		params = &CreateChannelFlowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateChannelFlow", params, optFns, c.addOperationCreateChannelFlowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateChannelFlowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateChannelFlowInput struct {

	// The ARN of the channel flow request.
	//
	// This member is required.
	AppInstanceArn *string

	// The client token for the request. An Idempotency token.
	//
	// This member is required.
	ClientRequestToken *string

	// The name of the channel flow.
	//
	// This member is required.
	Name *string

	// Information about the processor Lambda functions.
	//
	// This member is required.
	Processors []types.Processor

	// The tags for the creation request.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateChannelFlowOutput struct {

	// The ARN of the channel flow.
	ChannelFlowArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateChannelFlowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateChannelFlow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateChannelFlow{}, middleware.After)
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
	if err = addOpCreateChannelFlowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateChannelFlow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateChannelFlow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateChannelFlow",
	}
}
