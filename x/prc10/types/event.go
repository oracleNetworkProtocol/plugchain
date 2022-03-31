package types

const (
	EventTypeEditToken          = TypeMsgEditToken
	EventTypeBurnToken          = TypeMsgBurnToken
	EventTypeMintToken          = TypeMsgMintToken
	EventTypeIssueToken         = TypeMsgIssueToken
	EventTypeTransferOwnerToken = TypeMsgTransferOwnerToken

	AttributeKeySymbol        = "symbol"
	AttributeKeyOwner         = "owner"
	AttributeKeyScale         = "scale"
	AttributeKeyAmount        = "amount"
	AttributeKeyTo            = "to"
	AttributeKeyInitialSupply = "initial_supply"
)
