// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes an existing pool. Deleting a pool disassociates all origination
// identities from that pool. If the pool status isn't active or if deletion
// protection is enabled, an Error is returned. A pool is a collection of phone
// numbers and SenderIds. A pool can include one or more phone numbers and
// SenderIds that are associated with your Amazon Web Services account.
func (c *Client) DeletePool(ctx context.Context, params *DeletePoolInput, optFns ...func(*Options)) (*DeletePoolOutput, error) {
	if params == nil {
		params = &DeletePoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePool", params, optFns, c.addOperationDeletePoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePoolInput struct {

	// The PoolId or PoolArn of the pool to delete. You can use DescribePools to find
	// the values for PoolId and PoolArn .
	//
	// This member is required.
	PoolId *string

	noSmithyDocumentSerde
}

type DeletePoolOutput struct {

	// The time when the pool was created, in UNIX epoch time
	// (https://www.epochconverter.com/) format.
	CreatedTimestamp *time.Time

	// The message type that was associated with the deleted pool.
	MessageType types.MessageType

	// The name of the OptOutList that was associated with the deleted pool.
	OptOutListName *string

	// The Amazon Resource Name (ARN) of the pool that was deleted.
	PoolArn *string

	// The PoolId of the pool that was deleted.
	PoolId *string

	// By default this is set to false. When an end recipient sends a message that
	// begins with HELP or STOP to one of your dedicated numbers, Amazon Pinpoint
	// automatically replies with a customizable message and adds the end recipient to
	// the OptOutList. When set to true you're responsible for responding to HELP and
	// STOP requests. You're also responsible for tracking and honoring opt-out
	// requests.
	SelfManagedOptOutsEnabled bool

	// Indicates whether shared routes are enabled for the pool.
	SharedRoutesEnabled bool

	// The current status of the pool.
	//
	// * CREATING: The pool is currently being created
	// and isn't yet available for use.
	//
	// * ACTIVE: The pool is active and available for
	// use.
	//
	// * DELETING: The pool is being deleted.
	Status types.PoolStatus

	// The Amazon Resource Name (ARN) of the TwoWayChannel.
	TwoWayChannelArn *string

	// By default this is set to false. When set to true you can receive incoming text
	// messages from your end recipients.
	TwoWayEnabled bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeletePoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeletePool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeletePool{}, middleware.After)
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
	if err = addOpDeletePoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePool(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeletePool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "DeletePool",
	}
}
