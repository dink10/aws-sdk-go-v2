// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used by administrators to choose which groups in the directory should have
// access to upload and download files over the enabled protocols using Transfer
// Family. For example, a Microsoft Active Directory might contain 50,000 users,
// but only a small fraction might need the ability to transfer files to the
// server. An administrator can use CreateAccess to limit the access to the
// correct set of users who need this ability.
func (c *Client) CreateAccess(ctx context.Context, params *CreateAccessInput, optFns ...func(*Options)) (*CreateAccessOutput, error) {
	if params == nil {
		params = &CreateAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccess", params, optFns, c.addOperationCreateAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccessInput struct {

	// A unique identifier that is required to identify specific groups within your
	// directory. The users of the group that you associate have access to your Amazon
	// S3 or Amazon EFS resources over the enabled protocols using Transfer Family. If
	// you know the group name, you can view the SID values by running the following
	// command using Windows PowerShell. Get-ADGroup -Filter {samAccountName -like
	// "YourGroupName*"} -Properties * | Select SamAccountName,ObjectSid In that
	// command, replace YourGroupName with the name of your Active Directory group. The
	// regular expression used to validate this parameter is a string of characters
	// consisting of uppercase and lowercase alphanumeric characters with no spaces.
	// You can also include underscores or any of the following characters: =,.@:/-
	//
	// This member is required.
	ExternalId *string

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// that controls your users' access to your Amazon S3 bucket or Amazon EFS file
	// system. The policies attached to this role determine the level of access that
	// you want to provide your users when transferring files into and out of your
	// Amazon S3 bucket or Amazon EFS file system. The IAM role should also contain a
	// trust relationship that allows the server to access your resources when
	// servicing your users' transfer requests.
	//
	// This member is required.
	Role *string

	// A system-assigned unique identifier for a server instance. This is the specific
	// server that you added your user to.
	//
	// This member is required.
	ServerId *string

	// The landing directory (folder) for a user when they log in to the server using
	// the client. A HomeDirectory example is /bucket_name/home/mydirectory . The
	// HomeDirectory parameter is only used if HomeDirectoryType is set to LOGICAL .
	HomeDirectory *string

	// Logical directory mappings that specify what Amazon S3 or Amazon EFS paths and
	// keys should be visible to your user and how you want to make them visible. You
	// must specify the Entry and Target pair, where Entry shows how the path is made
	// visible and Target is the actual Amazon S3 or Amazon EFS path. If you only
	// specify a target, it is displayed as is. You also must ensure that your Identity
	// and Access Management (IAM) role provides access to paths in Target . This value
	// can be set only when HomeDirectoryType is set to LOGICAL. The following is an
	// Entry and Target pair example. [ { "Entry": "/directory1", "Target":
	// "/bucket_name/home/mydirectory" } ] In most cases, you can use this value
	// instead of the session policy to lock down your user to the designated home
	// directory (" chroot "). To do this, you can set Entry to / and set Target to
	// the HomeDirectory parameter value. The following is an Entry and Target pair
	// example for chroot . [ { "Entry": "/", "Target":
	// "/bucket_name/home/mydirectory" } ]
	HomeDirectoryMappings []types.HomeDirectoryMapEntry

	// The type of landing directory (folder) that you want your users' home directory
	// to be when they log in to the server. If you set it to PATH , the user will see
	// the absolute Amazon S3 bucket or Amazon EFS path as is in their file transfer
	// protocol clients. If you set it to LOGICAL , you need to provide mappings in the
	// HomeDirectoryMappings for how you want to make Amazon S3 or Amazon EFS paths
	// visible to your users. If HomeDirectoryType is LOGICAL , you must provide
	// mappings, using the HomeDirectoryMappings parameter. If, on the other hand,
	// HomeDirectoryType is PATH , you provide an absolute path using the HomeDirectory
	// parameter. You cannot have both HomeDirectory and HomeDirectoryMappings in your
	// template.
	HomeDirectoryType types.HomeDirectoryType

	// A session policy for your user so that you can use the same Identity and Access
	// Management (IAM) role across multiple users. This policy scopes down a user's
	// access to portions of their Amazon S3 bucket. Variables that you can use inside
	// this policy include ${Transfer:UserName} , ${Transfer:HomeDirectory} , and
	// ${Transfer:HomeBucket} . This policy applies only when the domain of ServerId
	// is Amazon S3. Amazon EFS does not use session policies. For session policies,
	// Transfer Family stores the policy as a JSON blob, instead of the Amazon Resource
	// Name (ARN) of the policy. You save the policy as a JSON blob and pass it in the
	// Policy argument. For an example of a session policy, see Example session policy (https://docs.aws.amazon.com/transfer/latest/userguide/session-policy.html)
	// . For more information, see AssumeRole (https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html)
	// in the Security Token Service API Reference.
	Policy *string

	// The full POSIX identity, including user ID ( Uid ), group ID ( Gid ), and any
	// secondary groups IDs ( SecondaryGids ), that controls your users' access to your
	// Amazon EFS file systems. The POSIX permissions that are set on files and
	// directories in your file system determine the level of access your users get
	// when transferring files into and out of your Amazon EFS file systems.
	PosixProfile *types.PosixProfile

	noSmithyDocumentSerde
}

type CreateAccessOutput struct {

	// The external identifier of the group whose users have access to your Amazon S3
	// or Amazon EFS resources over the enabled protocols using Transfer Family.
	//
	// This member is required.
	ExternalId *string

	// The identifier of the server that the user is attached to.
	//
	// This member is required.
	ServerId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAccess{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addCreateAccessResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccess(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "CreateAccess",
	}
}

type opCreateAccessResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateAccessResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateAccessResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "transfer"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "transfer"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("transfer")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addCreateAccessResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateAccessResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
