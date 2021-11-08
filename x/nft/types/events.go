package types

const (
	EventTypeIssueClass    = TypeMsgIssueClass
	EventTypeIssueNFT      = TypeMsgIssueNFT
	EventTypeEditNFT       = TypeMsgEditNFT
	EventTypeBurnNFT       = TypeMsgBurnNFT
	EventTypeTransferNFT   = TypeMsgTransferNFT
	EventTypeTransferClass = TypeMsgTransferClass

	AttributeKeyClassID   = "class_id"
	AttributeKeyNFTID     = "nft_id"
	AttributeKeyOwner     = "owner"
	AttributeKeyRecipient = "recipient"
)
