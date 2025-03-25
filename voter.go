//go:build !noAndroidStuff
// +build !noAndroidStuff

package identity

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rarimo/rarime-mobile-identity-sdk/contracts"
	"golang.org/x/crypto/sha3"
	"math/big"
)

type ProposalInfoDetailsConfigJSON struct {
	StartTimestamp      uint64              `json:"start_timestamp"`
	Duration            uint64              `json:"duration"`
	Multichoice         string              `json:"multichoice"`
	AcceptedOptions     []string            `json:"accepted_options"`
	Description         string              `json:"description"`
	VotingWhitelist     []common.Address    `json:"voting_whitelist"`
	VotingWhitelistData VotingWhitelistData `json:"voting_whitelist_data"`
}

type ProposalInfoDetailsJSON struct {
	ProposalSMT   common.Address                `json:"proposal_smt"`
	Status        uint8                         `json:"status"`
	Config        ProposalInfoDetailsConfigJSON `json:"config"`
	VotingResults [][]string                    `json:"voting_results"`
}

type ProposalInfoJSON struct {
	ProposalInfo    ProposalInfoDetailsJSON `json:"proposal_info"`
	ProposalEventId string                  `json:"proposal_event_id"`
}

type VotingWhitelistData struct {
	CitizenshipWhitelist                []uint64 `json:"citizenshipWhitelist"`
	IdentityCreationTimestampUpperBound uint64   `json:"identityCreationTimestampUpperBound"`
	IdentityCounterUpperBound           uint64   `json:"identityCounterUpperBound"`
	BirthDateUpperbound                 uint64   `json:"birthDateUpperbound"`
	ExpirationDateLowerBound            uint64   `json:"expirationDateLowerBound"`
}

type VoteResult struct {
	QuestionIndex int `json:"questionIndex"`
	AnswerIndex   int `json:"answerIndex,omitempty"`
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

	whilteListDataAbiTypes, _ := abi.NewType("tuple", "struct ProposalRules", []abi.ArgumentMarshaling{
		{
			Name: "citizenshipWhitelist",
			Type: "uint256[]",
		},
		{
			Name: "identityCreationTimestampUpperBound",
			Type: "uint256",
		},
		{
			Name: "identityCounterUpperBound",
			Type: "uint256",
		},
		{
			Name: "birthDateUpperbound",
			Type: "uint256",
		},
		{
			Name: "expirationDateLowerBound",
			Type: "uint256",
		},
	})

	whilteListDataAbiArgs := abi.Arguments{
		{
			Type: whilteListDataAbiTypes,
		},
	}

	whilteListDataAbiValues, err := whilteListDataAbiArgs.Unpack(info.Config.VotingWhitelistData[0])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack proposal rules: %v", err)
	}

	whilteListDataAbiValuesJSON, err := json.Marshal(whilteListDataAbiValues[0])
	if err != nil {
		return nil, fmt.Errorf("failed to marshal proposal rules: %v", err)
	}

	println()
	println(string(whilteListDataAbiValuesJSON))
	println()

	var votingWhitelistData VotingWhitelistData
	if err := json.Unmarshal(whilteListDataAbiValuesJSON, &votingWhitelistData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal VotingWhitelistData: %v", err)
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
				VotingWhitelistData: votingWhitelistData,
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

// CalculateVotingEventData calculates the voting event data.
func CalculateVotingEventData(voteJson []byte) ([]byte, error) {
	var voteResults []VoteResult
	if err := json.Unmarshal(voteJson, &voteResults); err != nil {
		return nil, fmt.Errorf("error unmarshalling vote: %v", err)
	}

	return TransformVote(voteResults)
}

// TransformVote transforms the vote.
func TransformVote(vote []VoteResult) ([]byte, error) {
	uint256Type, _ := abi.NewType("uint256[]", "", nil)

	args := abi.Arguments{
		{
			Type: uint256Type,
		},
	}

	var values []*big.Int
	for _, v := range vote {
		value := makeByteWithBits(v.AnswerIndex)

		values = append(values, big.NewInt(int64(value)))
	}

	packedVote, err := args.Pack(values)
	if err != nil {
		return nil, fmt.Errorf("failed to pack vote: %v", err)
	}

	hash := sha3.NewLegacyKeccak256()
	hash.Write(packedVote)
	hashBytes := hash.Sum(nil)
	uint248Bytes := hashBytes[1:32]

	return uint248Bytes, nil
}

func makeByteWithBits(n int) byte {
	if n < 0 || n > 8 {
		panic("n must be between 0 and 8")
	}

	return byte((1 << n))
}
