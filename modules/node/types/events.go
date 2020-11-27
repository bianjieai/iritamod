package types

// validator module event types
const (
	EventTypeCreateValidator = "create_validator"
	EventTypeUpdateValidator = "update_validator"
	EventTypeRemoveValidator = "remove_validator"
	EventTypeAddNode         = "add_node"
	EventTypeRemoveNode      = "remove_node"

	AttributeValueCategory = ModuleName
	AttributeKeyValidator  = "validator"
	AttributeKeyPubkey     = "pubkey"
	AttributeKeyID         = "id"
)
