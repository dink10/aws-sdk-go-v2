// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateModelCustomizationJob struct {
}

func (*validateOpCreateModelCustomizationJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateModelCustomizationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateModelCustomizationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateModelCustomizationJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateProvisionedModelThroughput struct {
}

func (*validateOpCreateProvisionedModelThroughput) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateProvisionedModelThroughput) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateProvisionedModelThroughputInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateProvisionedModelThroughputInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteCustomModel struct {
}

func (*validateOpDeleteCustomModel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteCustomModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteCustomModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteCustomModelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteProvisionedModelThroughput struct {
}

func (*validateOpDeleteProvisionedModelThroughput) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteProvisionedModelThroughput) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteProvisionedModelThroughputInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteProvisionedModelThroughputInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCustomModel struct {
}

func (*validateOpGetCustomModel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCustomModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCustomModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCustomModelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetFoundationModel struct {
}

func (*validateOpGetFoundationModel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetFoundationModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetFoundationModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetFoundationModelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetModelCustomizationJob struct {
}

func (*validateOpGetModelCustomizationJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetModelCustomizationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetModelCustomizationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetModelCustomizationJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetProvisionedModelThroughput struct {
}

func (*validateOpGetProvisionedModelThroughput) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetProvisionedModelThroughput) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetProvisionedModelThroughputInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetProvisionedModelThroughputInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutModelInvocationLoggingConfiguration struct {
}

func (*validateOpPutModelInvocationLoggingConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutModelInvocationLoggingConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutModelInvocationLoggingConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutModelInvocationLoggingConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopModelCustomizationJob struct {
}

func (*validateOpStopModelCustomizationJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopModelCustomizationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopModelCustomizationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopModelCustomizationJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateProvisionedModelThroughput struct {
}

func (*validateOpUpdateProvisionedModelThroughput) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateProvisionedModelThroughput) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateProvisionedModelThroughputInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateProvisionedModelThroughputInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateModelCustomizationJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateModelCustomizationJob{}, middleware.After)
}

func addOpCreateProvisionedModelThroughputValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateProvisionedModelThroughput{}, middleware.After)
}

func addOpDeleteCustomModelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteCustomModel{}, middleware.After)
}

func addOpDeleteProvisionedModelThroughputValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteProvisionedModelThroughput{}, middleware.After)
}

func addOpGetCustomModelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCustomModel{}, middleware.After)
}

func addOpGetFoundationModelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetFoundationModel{}, middleware.After)
}

func addOpGetModelCustomizationJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetModelCustomizationJob{}, middleware.After)
}

func addOpGetProvisionedModelThroughputValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetProvisionedModelThroughput{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPutModelInvocationLoggingConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutModelInvocationLoggingConfiguration{}, middleware.After)
}

func addOpStopModelCustomizationJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopModelCustomizationJob{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateProvisionedModelThroughputValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateProvisionedModelThroughput{}, middleware.After)
}

func validateCloudWatchConfig(v *types.CloudWatchConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CloudWatchConfig"}
	if v.LogGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LogGroupName"))
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if v.LargeDataDeliveryS3Config != nil {
		if err := validateS3Config(v.LargeDataDeliveryS3Config); err != nil {
			invalidParams.AddNested("LargeDataDeliveryS3Config", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLoggingConfig(v *types.LoggingConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LoggingConfig"}
	if v.CloudWatchConfig != nil {
		if err := validateCloudWatchConfig(v.CloudWatchConfig); err != nil {
			invalidParams.AddNested("CloudWatchConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3Config != nil {
		if err := validateS3Config(v.S3Config); err != nil {
			invalidParams.AddNested("S3Config", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOutputDataConfig(v *types.OutputDataConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OutputDataConfig"}
	if v.S3Uri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Uri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3Config(v *types.S3Config) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3Config"}
	if v.BucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BucketName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTrainingDataConfig(v *types.TrainingDataConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TrainingDataConfig"}
	if v.S3Uri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Uri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateValidationDataConfig(v *types.ValidationDataConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ValidationDataConfig"}
	if v.Validators == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Validators"))
	} else if v.Validators != nil {
		if err := validateValidators(v.Validators); err != nil {
			invalidParams.AddNested("Validators", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateValidator(v *types.Validator) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Validator"}
	if v.S3Uri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Uri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateValidators(v []types.Validator) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Validators"}
	for i := range v {
		if err := validateValidator(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateVpcConfig(v *types.VpcConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VpcConfig"}
	if v.SubnetIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubnetIds"))
	}
	if v.SecurityGroupIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecurityGroupIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateModelCustomizationJobInput(v *CreateModelCustomizationJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateModelCustomizationJobInput"}
	if v.JobName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobName"))
	}
	if v.CustomModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CustomModelName"))
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if v.BaseModelIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BaseModelIdentifier"))
	}
	if v.JobTags != nil {
		if err := validateTagList(v.JobTags); err != nil {
			invalidParams.AddNested("JobTags", err.(smithy.InvalidParamsError))
		}
	}
	if v.CustomModelTags != nil {
		if err := validateTagList(v.CustomModelTags); err != nil {
			invalidParams.AddNested("CustomModelTags", err.(smithy.InvalidParamsError))
		}
	}
	if v.TrainingDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TrainingDataConfig"))
	} else if v.TrainingDataConfig != nil {
		if err := validateTrainingDataConfig(v.TrainingDataConfig); err != nil {
			invalidParams.AddNested("TrainingDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.ValidationDataConfig != nil {
		if err := validateValidationDataConfig(v.ValidationDataConfig); err != nil {
			invalidParams.AddNested("ValidationDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputDataConfig"))
	} else if v.OutputDataConfig != nil {
		if err := validateOutputDataConfig(v.OutputDataConfig); err != nil {
			invalidParams.AddNested("OutputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.HyperParameters == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HyperParameters"))
	}
	if v.VpcConfig != nil {
		if err := validateVpcConfig(v.VpcConfig); err != nil {
			invalidParams.AddNested("VpcConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateProvisionedModelThroughputInput(v *CreateProvisionedModelThroughputInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateProvisionedModelThroughputInput"}
	if v.ModelUnits == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ModelUnits"))
	}
	if v.ProvisionedModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProvisionedModelName"))
	}
	if v.ModelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ModelId"))
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteCustomModelInput(v *DeleteCustomModelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteCustomModelInput"}
	if v.ModelIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ModelIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteProvisionedModelThroughputInput(v *DeleteProvisionedModelThroughputInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteProvisionedModelThroughputInput"}
	if v.ProvisionedModelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProvisionedModelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCustomModelInput(v *GetCustomModelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCustomModelInput"}
	if v.ModelIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ModelIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetFoundationModelInput(v *GetFoundationModelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetFoundationModelInput"}
	if v.ModelIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ModelIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetModelCustomizationJobInput(v *GetModelCustomizationJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetModelCustomizationJobInput"}
	if v.JobIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetProvisionedModelThroughputInput(v *GetProvisionedModelThroughputInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetProvisionedModelThroughputInput"}
	if v.ProvisionedModelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProvisionedModelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutModelInvocationLoggingConfigurationInput(v *PutModelInvocationLoggingConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutModelInvocationLoggingConfigurationInput"}
	if v.LoggingConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoggingConfig"))
	} else if v.LoggingConfig != nil {
		if err := validateLoggingConfig(v.LoggingConfig); err != nil {
			invalidParams.AddNested("LoggingConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopModelCustomizationJobInput(v *StopModelCustomizationJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopModelCustomizationJobInput"}
	if v.JobIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateProvisionedModelThroughputInput(v *UpdateProvisionedModelThroughputInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateProvisionedModelThroughputInput"}
	if v.ProvisionedModelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProvisionedModelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
