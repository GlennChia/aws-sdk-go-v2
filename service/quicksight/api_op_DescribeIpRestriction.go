// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a summary and status of IP Rules.
func (c *Client) DescribeIpRestriction(ctx context.Context, params *DescribeIpRestrictionInput, optFns ...func(*Options)) (*DescribeIpRestrictionOutput, error) {
	if params == nil {
		params = &DescribeIpRestrictionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeIpRestriction", params, optFns, c.addOperationDescribeIpRestrictionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeIpRestrictionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIpRestrictionInput struct {

	// Your AWS account ID.
	//
	// This member is required.
	AwsAccountId *string

	noSmithyDocumentSerde
}

type DescribeIpRestrictionOutput struct {

	// Your AWS account ID.
	AwsAccountId *string

	// Whether or not IP rules are enabled.
	Enabled *bool

	// Describes the IP rules with CIDR range and description.
	IpRestrictionRuleMap map[string]string

	// The ID of the describe request.
	RequestId *string

	// The status of a set of IP restrictions. A successful request returns a 200
	// status code.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeIpRestrictionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeIpRestriction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeIpRestriction{}, middleware.After)
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
	if err = addOpDescribeIpRestrictionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIpRestriction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeIpRestriction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeIpRestriction",
	}
}
