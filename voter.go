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

type ProposalInfoDetailsConfigJSON struct {
	StartTimestamp      uint64           `json:"startTimestamp"`
	Duration            uint64           `json:"duration"`
	Multichoice         string           `json:"multichoice"`
	AcceptedOptions     []string         `json:"acceptedOptions"`
	Description         string           `json:"Description"`
	VotingWhitelist     []common.Address `json:"voting_whitelist"`
	VotingWhitelistData [][]byte         `json:"voting_whitelist"`
}

type ProposalInfoDetailsJSON struct {
	ProposalSMT   common.Address                `json:"proposalSMT"`
	Status        uint8                         `json:"status"`
	Config        ProposalInfoDetailsConfigJSON `json:"config"`
	VotingResults [][]string                    `json:"voting_results"`
}

type ProposalInfoJSON struct {
	ProposalInfo    ProposalInfoDetailsJSON `json:"proposalInfo"`
	ProposalEventId string                  `json:"proposalEventId"`
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

	proposalInfo := &ProposalInfoJSON{
		ProposalInfo: ProposalInfoDetailsJSON{
			ProposalSMT: info.ProposalSMT,
			Status:      info.Status,
			Config: ProposalInfoDetailsConfigJSON{
				StartTimestamp:      info.Config.StartTimestamp,
				Duration:            info.Config.Duration,
				Multichoice:         info.Config.Multichoice.String(),
				AcceptedOptions:     convertBigIntSliceToStringSlice(info.Config.AcceptedOptions),
				Description:         info.Config.Description,
				VotingWhitelist:     info.Config.VotingWhitelist,
				VotingWhitelistData: info.Config.VotingWhitelistData,
			},
			VotingResults: convertVotingResults(info.VotingResults),
		},
		ProposalEventId: proposalEventId.String(),
	}

	proposalInfoJSON, err := json.Marshal(proposalInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json %v", err)
	}

	return proposalInfoJSON, nil
}

func convertBigIntSliceToStringSlice(options []*big.Int) []string {
	var result []string
	for _, opt := range options {
		if opt != nil {
			result = append(result, opt.String())
		} else {
			result = append(result, "")
		}
	}
	return result
}

func convertVotingResults(results [][8]*big.Int) [][]string {
	var finalResults [][]string
	for _, result := range results {
		row := make([]string, len(result))
		for i, val := range result {
			if val != nil {
				row[i] = val.String()
			} else {
				row[i] = ""
			}
		}
		finalResults = append(finalResults, row)
	}
	return finalResults
}
