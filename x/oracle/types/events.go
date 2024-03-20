package types

// claim module events
const (
	AttributeValueCategory = "oracle"
	AttributeKeyClaimHash  = "claim_hash"

	EventTypeDelegate     = "delegate"
	AttributeKeyDelegate  = "delegate"
	AttributeKeyValidator = "validator"

	EventTypeVote           = "vote"
	EventTypePrevote        = "prevote"
	AttributeKeyPrevoteHash = "prevote_hash"
)
