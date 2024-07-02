package types

// node module event types
const (
	EventTypeCreateValidator = "create_validator"
	EventTypeUpdateValidator = "update_validator"
	EventTypeRemoveValidator = "remove_validator"
	EventTypeGrantNode       = "grant_node"
	EventTypeRevokeNode      = "revoke_node"

	AttributeValueCategory = ModuleName
	AttributeKeyValidator  = "validator"
	AttributeKeyPubkey     = "pubkey"
	AttributeKeyID         = "id"
)
