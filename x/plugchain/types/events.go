package types

//Event types for the plugchain module.
const (
	EventTypeCreateToken = TypeMsgCreateToken
	EventTypeBurnToken   = TypeMsgBurnToken
)

// event value attri for the plugchain module
const (
	AttributeValueTokenSymbol         = "symbol"
	AttributeValueTokenOriginalSymbol = "original_symbol"
	AttributeValueTokenWholeName      = "whole_name"
	AttributeValueTokenOwner          = "owner"
	AttributeValueTokenTotalSupply    = "total_supply"
	AttributeValueTokenDescription    = "description"
	AttributeValueTokenMintable       = "mintable"
	AttributeValueTokenDecimal        = "decimal"

	AttributeValueTokenBurnAccount = "account"
)
