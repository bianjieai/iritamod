package types

// validator module event types
const (
	EventTypeCreateValidator = "create_validator"
	EventTypeUpdateValidator = "update_validator"
	EventTypeRemoveValidator = "remove_validator"

	AttributeValueCategory = ModuleName
	AttributeKeyValidator  = "validator"
	AttributeKeyPubkey     = "pubkey"
)
