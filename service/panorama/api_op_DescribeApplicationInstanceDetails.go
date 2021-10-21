// Code generated by smithy-go-codegen DO NOT EDIT.

package panorama

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/panorama/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about an application instance's configuration manifest.
func (c *Client) DescribeApplicationInstanceDetails(ctx context.Context, params *DescribeApplicationInstanceDetailsInput, optFns ...func(*Options)) (*DescribeApplicationInstanceDetailsOutput, error) {
	if params == nil {
		params = &DescribeApplicationInstanceDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeApplicationInstanceDetails", params, optFns, c.addOperationDescribeApplicationInstanceDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeApplicationInstanceDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeApplicationInstanceDetailsInput struct {

	// The application instance's ID.
	//
	// This member is required.
	ApplicationInstanceId *string

	noSmithyDocumentSerde
}

type DescribeApplicationInstanceDetailsOutput struct {

	// The application instance's ID.
	ApplicationInstanceId *string

	// The ID of the application instance that this instance replaced.
	ApplicationInstanceIdToReplace *string

	// When the application instance was created.
	CreatedTime *time.Time

	// The application instance's default runtime context device.
	DefaultRuntimeContextDevice *string

	// The application instance's description.
	Description *string

	// Parameter overrides for the configuration manifest.
	ManifestOverridesPayload types.ManifestOverridesPayload

	// The application instance's configuration manifest.
	ManifestPayload types.ManifestPayload

	// The application instance's name.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeApplicationInstanceDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeApplicationInstanceDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeApplicationInstanceDetails{}, middleware.After)
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
	if err = addOpDescribeApplicationInstanceDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApplicationInstanceDetails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeApplicationInstanceDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "panorama",
		OperationName: "DescribeApplicationInstanceDetails",
	}
}
