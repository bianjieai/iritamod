package types

// NewGenesisState constructs a new GenesisState instance
func NewGenesisState(identities []Identity) *GenesisState {
	return &GenesisState{Identities: identities}
}

// DefaultGenesisState gets the raw genesis raw message for testing
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}

// ValidateGenesis validates the provided identity genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	for _, identity := range data.Identities {
		if err := identity.Validate(); err != nil {
			return err
		}
	}

	return nil
}
