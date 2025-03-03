//go:build !noAndroidStuff
// +build !noAndroidStuff

package identity

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rarimo/rarime-mobile-identity-sdk/contracts"
	"math/big"
)

type ProposalInfo struct {
	proposalInfo    contracts.ProposalsStateProposalInfo
	proposalEventId big.Int
}

func GetProposalInfo(address string, rpc string, proposalId string) ([]byte, error) {
	conn, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Electroneum client: %v", err)
	}

	contract, err := contracts.NewProposalsState(common.HexToAddress(address), conn)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate Storage contract: %v", err)
	}

	proposalIdBN, ok := new(big.Int).SetString(proposalId, 10)
	if !ok {
		return nil, fmt.Errorf("error converting string to big int")
	}

	info, err := contract.GetProposalInfo(nil, proposalIdBN)
	if err != nil {
		return nil, fmt.Errorf("failed to GetProposalInfo contract: %v", err)
	}

	proposalEventId, err := contract.GetProposalEventId(nil, proposalIdBN)
	if err != nil {
		return nil, fmt.Errorf("failed to GetProposalEventId contract: %v", err)
	}

	proposalInfo := ProposalInfo{
		proposalInfo:    info,
		proposalEventId: *proposalEventId,
	}

	proposalInfoJSON, err := json.Marshal(proposalInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json %v", err)
	}

	return proposalInfoJSON, nil
}
