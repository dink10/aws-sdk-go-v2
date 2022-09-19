// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information for the specified import.
func (c *Client) GetImport(ctx context.Context, params *GetImportInput, optFns ...func(*Options)) (*GetImportOutput, error) {
	if params == nil {
		params = &GetImportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetImport", params, optFns, c.addOperationGetImportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetImportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetImportInput struct {

	// The ID for the import.
	//
	// This member is required.
	ImportId *string

	noSmithyDocumentSerde
}

type GetImportOutput struct {

	// The timestamp of the import's creation.
	CreatedTimestamp *time.Time

	// The destination event data store.
	Destinations []string

	// Used with StartEventTime to bound a StartImport request, and limit imported
	// trail events to only those events logged within a specified time period.
	EndEventTime *time.Time

	// The ID of the import.
	ImportId *string

	// The source S3 bucket.
	ImportSource *types.ImportSource

	// Provides statistics for the import.
	ImportStatistics *types.ImportStatistics

	// The status of the import.
	ImportStatus types.ImportStatus

	// Used with EndEventTime to bound a StartImport request, and limit imported trail
	// events to only those events logged within a specified time period.
	StartEventTime *time.Time

	// The timestamp of when the import was updated.
	UpdatedTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetImportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetImport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetImport{}, middleware.After)
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
	if err = addOpGetImportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetImport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetImport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "GetImport",
	}
}
