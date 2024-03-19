// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchainquery

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchainquery/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpBatchGetTokenBalance struct {
}

func (*validateOpBatchGetTokenBalance) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchGetTokenBalance) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchGetTokenBalanceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchGetTokenBalanceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAssetContract struct {
}

func (*validateOpGetAssetContract) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAssetContract) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAssetContractInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAssetContractInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTokenBalance struct {
}

func (*validateOpGetTokenBalance) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTokenBalance) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTokenBalanceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTokenBalanceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTransaction struct {
}

func (*validateOpGetTransaction) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTransaction) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTransactionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTransactionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListAssetContracts struct {
}

func (*validateOpListAssetContracts) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListAssetContracts) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListAssetContractsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListAssetContractsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListFilteredTransactionEvents struct {
}

func (*validateOpListFilteredTransactionEvents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListFilteredTransactionEvents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListFilteredTransactionEventsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListFilteredTransactionEventsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTokenBalances struct {
}

func (*validateOpListTokenBalances) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTokenBalances) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTokenBalancesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTokenBalancesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTransactionEvents struct {
}

func (*validateOpListTransactionEvents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTransactionEvents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTransactionEventsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTransactionEventsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTransactions struct {
}

func (*validateOpListTransactions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTransactions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTransactionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTransactionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpBatchGetTokenBalanceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchGetTokenBalance{}, middleware.After)
}

func addOpGetAssetContractValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAssetContract{}, middleware.After)
}

func addOpGetTokenBalanceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTokenBalance{}, middleware.After)
}

func addOpGetTransactionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTransaction{}, middleware.After)
}

func addOpListAssetContractsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListAssetContracts{}, middleware.After)
}

func addOpListFilteredTransactionEventsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListFilteredTransactionEvents{}, middleware.After)
}

func addOpListTokenBalancesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTokenBalances{}, middleware.After)
}

func addOpListTransactionEventsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTransactionEvents{}, middleware.After)
}

func addOpListTransactionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTransactions{}, middleware.After)
}

func validateAddressIdentifierFilter(v *types.AddressIdentifierFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AddressIdentifierFilter"}
	if v.TransactionEventToAddress == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TransactionEventToAddress"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateBatchGetTokenBalanceInputItem(v *types.BatchGetTokenBalanceInputItem) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchGetTokenBalanceInputItem"}
	if v.TokenIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TokenIdentifier"))
	} else if v.TokenIdentifier != nil {
		if err := validateTokenIdentifier(v.TokenIdentifier); err != nil {
			invalidParams.AddNested("TokenIdentifier", err.(smithy.InvalidParamsError))
		}
	}
	if v.OwnerIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OwnerIdentifier"))
	} else if v.OwnerIdentifier != nil {
		if err := validateOwnerIdentifier(v.OwnerIdentifier); err != nil {
			invalidParams.AddNested("OwnerIdentifier", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateConfirmationStatusFilter(v *types.ConfirmationStatusFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ConfirmationStatusFilter"}
	if v.Include == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Include"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateContractFilter(v *types.ContractFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ContractFilter"}
	if len(v.Network) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Network"))
	}
	if len(v.TokenStandard) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("TokenStandard"))
	}
	if v.DeployerAddress == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeployerAddress"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateContractIdentifier(v *types.ContractIdentifier) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ContractIdentifier"}
	if len(v.Network) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Network"))
	}
	if v.ContractAddress == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContractAddress"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateGetTokenBalanceInputList(v []types.BatchGetTokenBalanceInputItem) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTokenBalanceInputList"}
	for i := range v {
		if err := validateBatchGetTokenBalanceInputItem(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOwnerFilter(v *types.OwnerFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OwnerFilter"}
	if v.Address == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Address"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOwnerIdentifier(v *types.OwnerIdentifier) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OwnerIdentifier"}
	if v.Address == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Address"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTokenFilter(v *types.TokenFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TokenFilter"}
	if len(v.Network) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Network"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTokenIdentifier(v *types.TokenIdentifier) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TokenIdentifier"}
	if len(v.Network) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Network"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateVoutFilter(v *types.VoutFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VoutFilter"}
	if v.VoutSpent == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VoutSpent"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchGetTokenBalanceInput(v *BatchGetTokenBalanceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchGetTokenBalanceInput"}
	if v.GetTokenBalanceInputs != nil {
		if err := validateGetTokenBalanceInputList(v.GetTokenBalanceInputs); err != nil {
			invalidParams.AddNested("GetTokenBalanceInputs", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAssetContractInput(v *GetAssetContractInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAssetContractInput"}
	if v.ContractIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContractIdentifier"))
	} else if v.ContractIdentifier != nil {
		if err := validateContractIdentifier(v.ContractIdentifier); err != nil {
			invalidParams.AddNested("ContractIdentifier", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTokenBalanceInput(v *GetTokenBalanceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTokenBalanceInput"}
	if v.TokenIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TokenIdentifier"))
	} else if v.TokenIdentifier != nil {
		if err := validateTokenIdentifier(v.TokenIdentifier); err != nil {
			invalidParams.AddNested("TokenIdentifier", err.(smithy.InvalidParamsError))
		}
	}
	if v.OwnerIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OwnerIdentifier"))
	} else if v.OwnerIdentifier != nil {
		if err := validateOwnerIdentifier(v.OwnerIdentifier); err != nil {
			invalidParams.AddNested("OwnerIdentifier", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTransactionInput(v *GetTransactionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTransactionInput"}
	if v.TransactionHash == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TransactionHash"))
	}
	if len(v.Network) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Network"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListAssetContractsInput(v *ListAssetContractsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListAssetContractsInput"}
	if v.ContractFilter == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContractFilter"))
	} else if v.ContractFilter != nil {
		if err := validateContractFilter(v.ContractFilter); err != nil {
			invalidParams.AddNested("ContractFilter", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListFilteredTransactionEventsInput(v *ListFilteredTransactionEventsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListFilteredTransactionEventsInput"}
	if v.Network == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Network"))
	}
	if v.AddressIdentifierFilter == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AddressIdentifierFilter"))
	} else if v.AddressIdentifierFilter != nil {
		if err := validateAddressIdentifierFilter(v.AddressIdentifierFilter); err != nil {
			invalidParams.AddNested("AddressIdentifierFilter", err.(smithy.InvalidParamsError))
		}
	}
	if v.VoutFilter != nil {
		if err := validateVoutFilter(v.VoutFilter); err != nil {
			invalidParams.AddNested("VoutFilter", err.(smithy.InvalidParamsError))
		}
	}
	if v.ConfirmationStatusFilter != nil {
		if err := validateConfirmationStatusFilter(v.ConfirmationStatusFilter); err != nil {
			invalidParams.AddNested("ConfirmationStatusFilter", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTokenBalancesInput(v *ListTokenBalancesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTokenBalancesInput"}
	if v.OwnerFilter != nil {
		if err := validateOwnerFilter(v.OwnerFilter); err != nil {
			invalidParams.AddNested("OwnerFilter", err.(smithy.InvalidParamsError))
		}
	}
	if v.TokenFilter == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TokenFilter"))
	} else if v.TokenFilter != nil {
		if err := validateTokenFilter(v.TokenFilter); err != nil {
			invalidParams.AddNested("TokenFilter", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTransactionEventsInput(v *ListTransactionEventsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTransactionEventsInput"}
	if len(v.Network) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Network"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTransactionsInput(v *ListTransactionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTransactionsInput"}
	if v.Address == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Address"))
	}
	if len(v.Network) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Network"))
	}
	if v.ConfirmationStatusFilter != nil {
		if err := validateConfirmationStatusFilter(v.ConfirmationStatusFilter); err != nil {
			invalidParams.AddNested("ConfirmationStatusFilter", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
