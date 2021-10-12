package types

const (
	EventTypeIssueDenom    = TypeMsgIssueDenom
	EventTypeIssueNFT      = TypeMsgIssueNFT
	EventTypeEditNFT       = TypeMsgEditNFT
	EventTypeBurnNFT       = TypeMsgBurnNFT
	EventTypeTransferNFT   = TypeMsgTransferNFT
	EventTypeTransferDenom = TypeMsgTransferDenom

	AttributeKeyDenomID = "denom_id"
	AttributeKeyNFTID   = "nft_id"
	AttributeKeyOwner   = "owner"
)
