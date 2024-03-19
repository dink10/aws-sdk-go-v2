// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchainquery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchainquery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the transaction events for an address on the blockchain. This
// operation is only supported on the Bitcoin networks.
func (c *Client) ListFilteredTransactionEvents(ctx context.Context, params *ListFilteredTransactionEventsInput, optFns ...func(*Options)) (*ListFilteredTransactionEventsOutput, error) {
	if params == nil {
		params = &ListFilteredTransactionEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFilteredTransactionEvents", params, optFns, c.addOperationListFilteredTransactionEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFilteredTransactionEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFilteredTransactionEventsInput struct {

	// This is the unique public address on the blockchain for which the transaction
	// events are being requested.
	//
	// This member is required.
	AddressIdentifierFilter *types.AddressIdentifierFilter

	// The blockchain network where the transaction occurred. Valid Values:
	// BITCOIN_MAINNET | BITCOIN_TESTNET
	//
	// This member is required.
	Network *string

	// The container for the ConfirmationStatusFilter that filters for the  finality  (https://docs.aws.amazon.com/managed-blockchain/latest/ambq-dg/key-concepts.html#finality)
	// of the results.
	ConfirmationStatusFilter *types.ConfirmationStatusFilter

	// The maximum number of transaction events to list. Default: 100 Even if
	// additional results can be retrieved, the request can return less results than
	// maxResults or an empty array of results. To retrieve the next set of results,
	// make another request with the returned nextToken value. The value of nextToken
	// is null when there are no more results to return
	MaxResults *int32

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// The order by which the results will be sorted.
	Sort *types.ListFilteredTransactionEventsSort

	// This container specifies the time frame for the transaction events returned in
	// the response.
	TimeFilter *types.TimeFilter

	// This container specifies filtering attributes related to BITCOIN_VOUT event
	// types
	VoutFilter *types.VoutFilter

	noSmithyDocumentSerde
}

type ListFilteredTransactionEventsOutput struct {

	// The transaction events returned by the request.
	//
	// This member is required.
	Events []types.TransactionEvent

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFilteredTransactionEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFilteredTransactionEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFilteredTransactionEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFilteredTransactionEvents"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListFilteredTransactionEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFilteredTransactionEvents(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// ListFilteredTransactionEventsAPIClient is a client that implements the
// ListFilteredTransactionEvents operation.
type ListFilteredTransactionEventsAPIClient interface {
	ListFilteredTransactionEvents(context.Context, *ListFilteredTransactionEventsInput, ...func(*Options)) (*ListFilteredTransactionEventsOutput, error)
}

var _ ListFilteredTransactionEventsAPIClient = (*Client)(nil)

// ListFilteredTransactionEventsPaginatorOptions is the paginator options for
// ListFilteredTransactionEvents
type ListFilteredTransactionEventsPaginatorOptions struct {
	// The maximum number of transaction events to list. Default: 100 Even if
	// additional results can be retrieved, the request can return less results than
	// maxResults or an empty array of results. To retrieve the next set of results,
	// make another request with the returned nextToken value. The value of nextToken
	// is null when there are no more results to return
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFilteredTransactionEventsPaginator is a paginator for
// ListFilteredTransactionEvents
type ListFilteredTransactionEventsPaginator struct {
	options   ListFilteredTransactionEventsPaginatorOptions
	client    ListFilteredTransactionEventsAPIClient
	params    *ListFilteredTransactionEventsInput
	nextToken *string
	firstPage bool
}

// NewListFilteredTransactionEventsPaginator returns a new
// ListFilteredTransactionEventsPaginator
func NewListFilteredTransactionEventsPaginator(client ListFilteredTransactionEventsAPIClient, params *ListFilteredTransactionEventsInput, optFns ...func(*ListFilteredTransactionEventsPaginatorOptions)) *ListFilteredTransactionEventsPaginator {
	if params == nil {
		params = &ListFilteredTransactionEventsInput{}
	}

	options := ListFilteredTransactionEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFilteredTransactionEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFilteredTransactionEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFilteredTransactionEvents page.
func (p *ListFilteredTransactionEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFilteredTransactionEventsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListFilteredTransactionEvents(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListFilteredTransactionEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFilteredTransactionEvents",
	}
}
