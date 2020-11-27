package types

// admin module event types
const (
	EventTypeAddRoles       = "add_roles"
	EventTypeRemoveRoles    = "remove_roles"
	EventTypeBlockAccount   = "block_account"
	EventTypeUnblockAccount = "unblock_account"

	AttributeKeyAccount    = "account"
	AttributeValueCategory = ModuleName
)
