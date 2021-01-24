package types

// perm module event types
const (
	EventTypeAssignRoles    = "assign_roles"
	EventTypeUnassignRoles  = "unassign_roles"
	EventTypeBlockAccount   = "block_account"
	EventTypeUnblockAccount = "unblock_account"

	AttributeKeyAccount    = "account"
	AttributeValueCategory = ModuleName
)
