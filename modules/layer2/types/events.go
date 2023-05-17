package types

const (
	EventTypeCreateL2Space   = "create_l2_space"
	EventTypeTransferL2Space = "transfer_l2_space"
	EventTypeCreateL2Record  = "create_l2_record"

	AttributeKeySender       = "sender"
	AttributeKeyOwner        = "owner"
	AttributeKeyRecipient    = "recipient"
	AttributeKeySpaceId      = "space_id"
	AttributeKeyRecordHeight = "record_height"

	EventTypeCreateNFTs          = "create_nfts"
	EventTypeUpdateNFTs          = "update_nfts"
	EventTypeDeleteNFTs          = "delete_nfts"
	EventTypeUpdateClassesForNFT = "update_classes_for_nft"

	AttributeKeyClassIdForNFT       = "class_id_for_nft"
	AttributeKeyTokenIdForNFT       = "token_id_for_nft"
	AttributeKeyClassesAmountForNFT = "classes_amount_for_nft"
	AttributeKeyTokenAmountForNFT   = "token_amount_for_nft"

	EventTypeDepositClassForNFT  = "deposit_class_for_nft"
	EventTypeWithdrawClassForNFT = "withdraw_class_for_nft"
	EventTypeDepositTokenForNFT  = "deposit_token_for_nft"
	EventTypeWithdrawTokenForNFT = "withdraw_token_for_nft"
)
