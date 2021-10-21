// Code generated by smithy-go-codegen DO NOT EDIT.

package panorama

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/panorama/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of jobs.
func (c *Client) ListDevicesJobs(ctx context.Context, params *ListDevicesJobsInput, optFns ...func(*Options)) (*ListDevicesJobsOutput, error) {
	if params == nil {
		params = &ListDevicesJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDevicesJobs", params, optFns, c.addOperationListDevicesJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDevicesJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDevicesJobsInput struct {

	// Filter results by the job's target device ID.
	DeviceId *string

	// The maximum number of device jobs to return in one page of results.
	MaxResults int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDevicesJobsOutput struct {

	// A list of jobs.
	DeviceJobs []types.DeviceJob

	// A pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDevicesJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDevicesJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDevicesJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDevicesJobs(options.Region), middleware.Before); err != nil {
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

// ListDevicesJobsAPIClient is a client that implements the ListDevicesJobs
// operation.
type ListDevicesJobsAPIClient interface {
	ListDevicesJobs(context.Context, *ListDevicesJobsInput, ...func(*Options)) (*ListDevicesJobsOutput, error)
}

var _ ListDevicesJobsAPIClient = (*Client)(nil)

// ListDevicesJobsPaginatorOptions is the paginator options for ListDevicesJobs
type ListDevicesJobsPaginatorOptions struct {
	// The maximum number of device jobs to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDevicesJobsPaginator is a paginator for ListDevicesJobs
type ListDevicesJobsPaginator struct {
	options   ListDevicesJobsPaginatorOptions
	client    ListDevicesJobsAPIClient
	params    *ListDevicesJobsInput
	nextToken *string
	firstPage bool
}

// NewListDevicesJobsPaginator returns a new ListDevicesJobsPaginator
func NewListDevicesJobsPaginator(client ListDevicesJobsAPIClient, params *ListDevicesJobsInput, optFns ...func(*ListDevicesJobsPaginatorOptions)) *ListDevicesJobsPaginator {
	if params == nil {
		params = &ListDevicesJobsInput{}
	}

	options := ListDevicesJobsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDevicesJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDevicesJobsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDevicesJobs page.
func (p *ListDevicesJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDevicesJobsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListDevicesJobs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListDevicesJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "panorama",
		OperationName: "ListDevicesJobs",
	}
}
