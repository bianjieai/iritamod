package types

// NewGenesisState constructs a new GenesisState instance
func NewGenesisState(nodes []Node) *GenesisState {
	return &GenesisState{Nodes: nodes}
}

// DefaultGenesisState gets the raw genesis raw message for testing
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}

// ValidateGenesis validates the provided node genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	for _, node := range data.Nodes {
		if err := node.Validate(); err != nil {
			return err
		}
	}

	return nil
}
