package types

import (
	query "github.com/cosmos/cosmos-sdk/types/query"
)

// QueryAllClaimsParams querey params struct
type QueryAllClaimsParams struct {
	Page  int `json:"page" yaml:"page"`
	Limit int `json:"limit" yaml:"limit"`
}

// NewQueryAllClaimsParams creates a new instance of QueryAllClaimsParams.
func NewQueryAllClaimsParams(page, limit int) QueryAllClaimsParams {
	return QueryAllClaimsParams{Page: page, Limit: limit}
}

// NewQueryClaimRequest creates a new instance of QueryClaimRequest.
func NewQueryClaimRequest(hash string) *QueryClaimRequest {
	return &QueryClaimRequest{ClaimHash: hash}
}

// NewQueryAllClaimsRequest creates a new instance of QueryAllClaimRequest.
func NewQueryAllClaimsRequest(pageReq *query.PageRequest) *QueryAllClaimsRequest {
	return &QueryAllClaimsRequest{Pagination: pageReq}
}
