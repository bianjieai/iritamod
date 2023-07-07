package genutil

import (
	"encoding/json"
	"path/filepath"
	"time"

	cfg "github.com/cometbft/cometbft/config"
	tmos "github.com/cometbft/cometbft/libs/os"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
	ctmtypes "github.com/cometbft/cometbft/types"
)

// ExportGenesisFile creates and writes the genesis configuration to disk. An
// error is returned if building or writing the configuration to file fails.
func ExportGenesisFile(genDoc *ctmtypes.GenesisDoc, genFile string) error {
	if err := genDoc.ValidateAndComplete(); err != nil {
		return err
	}

	return genDoc.SaveAs(genFile)
}

// ExportGenesisFileWithTime creates and writes the genesis configuration to disk.
// An error is returned if building or writing the configuration to file fails.
func ExportGenesisFileWithTime(
	genFile, chainID string, validators []ctmtypes.GenesisValidator,
	appState json.RawMessage, genTime time.Time,
) error {
	genDoc := ctmtypes.GenesisDoc{
		GenesisTime: genTime,
		ChainID:     chainID,
		Validators:  validators,
		AppState:    appState,
	}

	if err := genDoc.ValidateAndComplete(); err != nil {
		return err
	}

	return genDoc.SaveAs(genFile)
}

// InitializeNodeValidatorFiles creates private validator and p2p configuration files.
func InitializeNodeValidatorFiles(
	config *cfg.Config,
) (nodeKey *p2p.NodeKey, pv *privval.FilePV, err error) {
	nodeKey, err = p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		return nil, nil, err
	}

	pvKeyFile := config.PrivValidatorKeyFile()
	if err := tmos.EnsureDir(filepath.Dir(pvKeyFile), 0777); err != nil {
		return nil, nil, err
	}

	pvStateFile := config.PrivValidatorStateFile()
	if err := tmos.EnsureDir(filepath.Dir(pvStateFile), 0777); err != nil {
		return nil, nil, err
	}

	return nodeKey, privval.LoadOrGenFilePV(pvKeyFile, pvStateFile), nil
}
