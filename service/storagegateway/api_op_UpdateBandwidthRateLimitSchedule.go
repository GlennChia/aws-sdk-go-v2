// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the bandwidth rate limit schedule for a specified gateway. By default,
// gateways do not have bandwidth rate limit schedules, which means no bandwidth
// rate limiting is in effect. Use this to initiate or update a gateway's bandwidth
// rate limit schedule. This operation is supported in the volume and tape gateway
// types.
func (c *Client) UpdateBandwidthRateLimitSchedule(ctx context.Context, params *UpdateBandwidthRateLimitScheduleInput, optFns ...func(*Options)) (*UpdateBandwidthRateLimitScheduleOutput, error) {
	if params == nil {
		params = &UpdateBandwidthRateLimitScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBandwidthRateLimitSchedule", params, optFns, c.addOperationUpdateBandwidthRateLimitScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBandwidthRateLimitScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateBandwidthRateLimitScheduleInput struct {

	// An array containing bandwidth rate limit schedule intervals for a gateway. When
	// no bandwidth rate limit intervals have been scheduled, the array is empty.
	//
	// This member is required.
	BandwidthRateLimitIntervals []types.BandwidthRateLimitInterval

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	noSmithyDocumentSerde
}

type UpdateBandwidthRateLimitScheduleOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and Amazon Web Services Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateBandwidthRateLimitScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateBandwidthRateLimitSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateBandwidthRateLimitSchedule{}, middleware.After)
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
	if err = addOpUpdateBandwidthRateLimitScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBandwidthRateLimitSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateBandwidthRateLimitSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "UpdateBandwidthRateLimitSchedule",
	}
}
