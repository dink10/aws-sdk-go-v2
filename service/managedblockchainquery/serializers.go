// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchainquery

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchainquery/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpBatchGetTokenBalance struct {
}

func (*awsRestjson1_serializeOpBatchGetTokenBalance) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpBatchGetTokenBalance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*BatchGetTokenBalanceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/batch-get-token-balance")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentBatchGetTokenBalanceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsBatchGetTokenBalanceInput(v *BatchGetTokenBalanceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentBatchGetTokenBalanceInput(v *BatchGetTokenBalanceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.GetTokenBalanceInputs != nil {
		ok := object.Key("getTokenBalanceInputs")
		if err := awsRestjson1_serializeDocumentGetTokenBalanceInputList(v.GetTokenBalanceInputs, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetAssetContract struct {
}

func (*awsRestjson1_serializeOpGetAssetContract) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetAssetContract) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetAssetContractInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/get-asset-contract")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentGetAssetContractInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetAssetContractInput(v *GetAssetContractInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentGetAssetContractInput(v *GetAssetContractInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ContractIdentifier != nil {
		ok := object.Key("contractIdentifier")
		if err := awsRestjson1_serializeDocumentContractIdentifier(v.ContractIdentifier, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetTokenBalance struct {
}

func (*awsRestjson1_serializeOpGetTokenBalance) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetTokenBalance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetTokenBalanceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/get-token-balance")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentGetTokenBalanceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetTokenBalanceInput(v *GetTokenBalanceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentGetTokenBalanceInput(v *GetTokenBalanceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AtBlockchainInstant != nil {
		ok := object.Key("atBlockchainInstant")
		if err := awsRestjson1_serializeDocumentBlockchainInstant(v.AtBlockchainInstant, ok); err != nil {
			return err
		}
	}

	if v.OwnerIdentifier != nil {
		ok := object.Key("ownerIdentifier")
		if err := awsRestjson1_serializeDocumentOwnerIdentifier(v.OwnerIdentifier, ok); err != nil {
			return err
		}
	}

	if v.TokenIdentifier != nil {
		ok := object.Key("tokenIdentifier")
		if err := awsRestjson1_serializeDocumentTokenIdentifier(v.TokenIdentifier, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetTransaction struct {
}

func (*awsRestjson1_serializeOpGetTransaction) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetTransaction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetTransactionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/get-transaction")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentGetTransactionInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetTransactionInput(v *GetTransactionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentGetTransactionInput(v *GetTransactionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.Network) > 0 {
		ok := object.Key("network")
		ok.String(string(v.Network))
	}

	if v.TransactionHash != nil {
		ok := object.Key("transactionHash")
		ok.String(*v.TransactionHash)
	}

	return nil
}

type awsRestjson1_serializeOpListAssetContracts struct {
}

func (*awsRestjson1_serializeOpListAssetContracts) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListAssetContracts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListAssetContractsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/list-asset-contracts")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentListAssetContractsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListAssetContractsInput(v *ListAssetContractsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentListAssetContractsInput(v *ListAssetContractsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ContractFilter != nil {
		ok := object.Key("contractFilter")
		if err := awsRestjson1_serializeDocumentContractFilter(v.ContractFilter, ok); err != nil {
			return err
		}
	}

	if v.MaxResults != nil {
		ok := object.Key("maxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

type awsRestjson1_serializeOpListFilteredTransactionEvents struct {
}

func (*awsRestjson1_serializeOpListFilteredTransactionEvents) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListFilteredTransactionEvents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListFilteredTransactionEventsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/list-filtered-transaction-events")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentListFilteredTransactionEventsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListFilteredTransactionEventsInput(v *ListFilteredTransactionEventsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentListFilteredTransactionEventsInput(v *ListFilteredTransactionEventsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AddressIdentifierFilter != nil {
		ok := object.Key("addressIdentifierFilter")
		if err := awsRestjson1_serializeDocumentAddressIdentifierFilter(v.AddressIdentifierFilter, ok); err != nil {
			return err
		}
	}

	if v.ConfirmationStatusFilter != nil {
		ok := object.Key("confirmationStatusFilter")
		if err := awsRestjson1_serializeDocumentConfirmationStatusFilter(v.ConfirmationStatusFilter, ok); err != nil {
			return err
		}
	}

	if v.MaxResults != nil {
		ok := object.Key("maxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.Network != nil {
		ok := object.Key("network")
		ok.String(*v.Network)
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	if v.Sort != nil {
		ok := object.Key("sort")
		if err := awsRestjson1_serializeDocumentListFilteredTransactionEventsSort(v.Sort, ok); err != nil {
			return err
		}
	}

	if v.TimeFilter != nil {
		ok := object.Key("timeFilter")
		if err := awsRestjson1_serializeDocumentTimeFilter(v.TimeFilter, ok); err != nil {
			return err
		}
	}

	if v.VoutFilter != nil {
		ok := object.Key("voutFilter")
		if err := awsRestjson1_serializeDocumentVoutFilter(v.VoutFilter, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListTokenBalances struct {
}

func (*awsRestjson1_serializeOpListTokenBalances) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListTokenBalances) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTokenBalancesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/list-token-balances")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentListTokenBalancesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListTokenBalancesInput(v *ListTokenBalancesInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentListTokenBalancesInput(v *ListTokenBalancesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("maxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	if v.OwnerFilter != nil {
		ok := object.Key("ownerFilter")
		if err := awsRestjson1_serializeDocumentOwnerFilter(v.OwnerFilter, ok); err != nil {
			return err
		}
	}

	if v.TokenFilter != nil {
		ok := object.Key("tokenFilter")
		if err := awsRestjson1_serializeDocumentTokenFilter(v.TokenFilter, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListTransactionEvents struct {
}

func (*awsRestjson1_serializeOpListTransactionEvents) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListTransactionEvents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTransactionEventsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/list-transaction-events")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentListTransactionEventsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListTransactionEventsInput(v *ListTransactionEventsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentListTransactionEventsInput(v *ListTransactionEventsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("maxResults")
		ok.Integer(*v.MaxResults)
	}

	if len(v.Network) > 0 {
		ok := object.Key("network")
		ok.String(string(v.Network))
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	if v.TransactionHash != nil {
		ok := object.Key("transactionHash")
		ok.String(*v.TransactionHash)
	}

	if v.TransactionId != nil {
		ok := object.Key("transactionId")
		ok.String(*v.TransactionId)
	}

	return nil
}

type awsRestjson1_serializeOpListTransactions struct {
}

func (*awsRestjson1_serializeOpListTransactions) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListTransactions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTransactionsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/list-transactions")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentListTransactionsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListTransactionsInput(v *ListTransactionsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentListTransactionsInput(v *ListTransactionsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Address != nil {
		ok := object.Key("address")
		ok.String(*v.Address)
	}

	if v.ConfirmationStatusFilter != nil {
		ok := object.Key("confirmationStatusFilter")
		if err := awsRestjson1_serializeDocumentConfirmationStatusFilter(v.ConfirmationStatusFilter, ok); err != nil {
			return err
		}
	}

	if v.FromBlockchainInstant != nil {
		ok := object.Key("fromBlockchainInstant")
		if err := awsRestjson1_serializeDocumentBlockchainInstant(v.FromBlockchainInstant, ok); err != nil {
			return err
		}
	}

	if v.MaxResults != nil {
		ok := object.Key("maxResults")
		ok.Integer(*v.MaxResults)
	}

	if len(v.Network) > 0 {
		ok := object.Key("network")
		ok.String(string(v.Network))
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	if v.Sort != nil {
		ok := object.Key("sort")
		if err := awsRestjson1_serializeDocumentListTransactionsSort(v.Sort, ok); err != nil {
			return err
		}
	}

	if v.ToBlockchainInstant != nil {
		ok := object.Key("toBlockchainInstant")
		if err := awsRestjson1_serializeDocumentBlockchainInstant(v.ToBlockchainInstant, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentAddressIdentifierFilter(v *types.AddressIdentifierFilter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.TransactionEventToAddress != nil {
		ok := object.Key("transactionEventToAddress")
		if err := awsRestjson1_serializeDocumentChainAddresses(v.TransactionEventToAddress, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentBatchGetTokenBalanceInputItem(v *types.BatchGetTokenBalanceInputItem, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AtBlockchainInstant != nil {
		ok := object.Key("atBlockchainInstant")
		if err := awsRestjson1_serializeDocumentBlockchainInstant(v.AtBlockchainInstant, ok); err != nil {
			return err
		}
	}

	if v.OwnerIdentifier != nil {
		ok := object.Key("ownerIdentifier")
		if err := awsRestjson1_serializeDocumentOwnerIdentifier(v.OwnerIdentifier, ok); err != nil {
			return err
		}
	}

	if v.TokenIdentifier != nil {
		ok := object.Key("tokenIdentifier")
		if err := awsRestjson1_serializeDocumentTokenIdentifier(v.TokenIdentifier, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentBlockchainInstant(v *types.BlockchainInstant, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Time != nil {
		ok := object.Key("time")
		ok.Double(smithytime.FormatEpochSeconds(*v.Time))
	}

	return nil
}

func awsRestjson1_serializeDocumentChainAddresses(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentConfirmationStatusFilter(v *types.ConfirmationStatusFilter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Include != nil {
		ok := object.Key("include")
		if err := awsRestjson1_serializeDocumentConfirmationStatusIncludeList(v.Include, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentConfirmationStatusIncludeList(v []types.ConfirmationStatus, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}

func awsRestjson1_serializeDocumentContractFilter(v *types.ContractFilter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DeployerAddress != nil {
		ok := object.Key("deployerAddress")
		ok.String(*v.DeployerAddress)
	}

	if len(v.Network) > 0 {
		ok := object.Key("network")
		ok.String(string(v.Network))
	}

	if len(v.TokenStandard) > 0 {
		ok := object.Key("tokenStandard")
		ok.String(string(v.TokenStandard))
	}

	return nil
}

func awsRestjson1_serializeDocumentContractIdentifier(v *types.ContractIdentifier, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ContractAddress != nil {
		ok := object.Key("contractAddress")
		ok.String(*v.ContractAddress)
	}

	if len(v.Network) > 0 {
		ok := object.Key("network")
		ok.String(string(v.Network))
	}

	return nil
}

func awsRestjson1_serializeDocumentGetTokenBalanceInputList(v []types.BatchGetTokenBalanceInputItem, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentBatchGetTokenBalanceInputItem(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentListFilteredTransactionEventsSort(v *types.ListFilteredTransactionEventsSort, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.SortBy) > 0 {
		ok := object.Key("sortBy")
		ok.String(string(v.SortBy))
	}

	if len(v.SortOrder) > 0 {
		ok := object.Key("sortOrder")
		ok.String(string(v.SortOrder))
	}

	return nil
}

func awsRestjson1_serializeDocumentListTransactionsSort(v *types.ListTransactionsSort, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.SortBy) > 0 {
		ok := object.Key("sortBy")
		ok.String(string(v.SortBy))
	}

	if len(v.SortOrder) > 0 {
		ok := object.Key("sortOrder")
		ok.String(string(v.SortOrder))
	}

	return nil
}

func awsRestjson1_serializeDocumentOwnerFilter(v *types.OwnerFilter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Address != nil {
		ok := object.Key("address")
		ok.String(*v.Address)
	}

	return nil
}

func awsRestjson1_serializeDocumentOwnerIdentifier(v *types.OwnerIdentifier, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Address != nil {
		ok := object.Key("address")
		ok.String(*v.Address)
	}

	return nil
}

func awsRestjson1_serializeDocumentTimeFilter(v *types.TimeFilter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.From != nil {
		ok := object.Key("from")
		if err := awsRestjson1_serializeDocumentBlockchainInstant(v.From, ok); err != nil {
			return err
		}
	}

	if v.To != nil {
		ok := object.Key("to")
		if err := awsRestjson1_serializeDocumentBlockchainInstant(v.To, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentTokenFilter(v *types.TokenFilter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ContractAddress != nil {
		ok := object.Key("contractAddress")
		ok.String(*v.ContractAddress)
	}

	if len(v.Network) > 0 {
		ok := object.Key("network")
		ok.String(string(v.Network))
	}

	if v.TokenId != nil {
		ok := object.Key("tokenId")
		ok.String(*v.TokenId)
	}

	return nil
}

func awsRestjson1_serializeDocumentTokenIdentifier(v *types.TokenIdentifier, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ContractAddress != nil {
		ok := object.Key("contractAddress")
		ok.String(*v.ContractAddress)
	}

	if len(v.Network) > 0 {
		ok := object.Key("network")
		ok.String(string(v.Network))
	}

	if v.TokenId != nil {
		ok := object.Key("tokenId")
		ok.String(*v.TokenId)
	}

	return nil
}

func awsRestjson1_serializeDocumentVoutFilter(v *types.VoutFilter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.VoutSpent != nil {
		ok := object.Key("voutSpent")
		ok.Boolean(*v.VoutSpent)
	}

	return nil
}
