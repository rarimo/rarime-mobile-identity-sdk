package identity

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ProposalInfoIndexed struct {
	ProposalInfo ProposalsStateProposalInfo
	Index        int
}

func GetStateInfosMulticall(proposalAddress string, rpc string, multicallAddress string, _startIndex string, _endIndex string) ([]byte, error) {

	startIndex, ok := new(big.Int).SetString(_startIndex, 10)
	if !ok {
		return nil, fmt.Errorf("failed to convert string to int")
	}

	endIndex, ok := new(big.Int).SetString(_endIndex, 10)
	if !ok {
		return nil, fmt.Errorf("failed to convert string to int")
	}

	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	// Create an instance of the Multicall contract
	multicall, err := NewMulticallStorage(common.HexToAddress(multicallAddress), client)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new Multicall instance: %v", err)
	}

	// Encode the call data
	proposalABI, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %v", err)
	}

	proposalInfoTyp, err := abi.NewType("tuple", "struct ProposalInfo", []abi.ArgumentMarshaling{
		{Name: "proposalSMT", Type: "address"},
		{Name: "status", Type: "uint8"},
		{Name: "config", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "startTimestamp", Type: "uint64"},
			{Name: "duration", Type: "uint64"},
			{Name: "multichoice", Type: "uint256"},
			{Name: "acceptedOptions", Type: "uint256[]"},
			{Name: "description", Type: "string"},
			{Name: "votingWhitelist", Type: "address[]"},
			{Name: "votingWhitelistData", Type: "bytes[]"},
		}},
		{Name: "votingResults", Type: "uint256[8][]"},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create proposalInfo type: %v", err)
	}

	args := abi.Arguments{
		{
			Type: proposalInfoTyp,
		},
	}

	// Prepare the Multicall
	var calls []Multicall3Call3

	for i := new(big.Int).Set(startIndex); i.Cmp(endIndex) <= 0; i.Add(i, big.NewInt(1)) {

		callDataBytes, err := proposalABI.Pack("getProposalInfo", i)
		if err != nil {
			return nil, fmt.Errorf("failed to pack call data: %v", err)
		}

		item := Multicall3Call3{
			Target:   common.HexToAddress(proposalAddress),
			CallData: callDataBytes,
		}
		calls = append(calls, item)
	}

	// Execute the multicall
	result, err := multicall.Aggregate3(&bind.CallOpts{}, calls)
	if err != nil {
		return nil, fmt.Errorf("failed to execute multicall: %v", err)
	}

	// Check and debug return data

	// Unpack the result
	//var proposalInfo identity.ProposalsStateProposalInfo

	var proposalInfos []ProposalsStateProposalInfo

	for _, res := range result {
		values, err := args.Unpack(res.ReturnData)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack Proposal Info: %v", err)
		}

		valuesJSON, err := json.Marshal(values[0])
		if err != nil {
			return nil, fmt.Errorf("failed to marshal Proposal Info: %v", err)
		}

		stateInfo := new(ProposalsStateProposalInfo)

		err = json.Unmarshal(valuesJSON, stateInfo)

		if err != nil {
			return nil, fmt.Errorf("failed to Unmarshal Proposal Infos: %v", err)
		}

		proposalInfos = append(proposalInfos, *stateInfo)

	}

	var proposalInfoIndexed []ProposalInfoIndexed
	var index int = 0

	for i := new(big.Int).Set(startIndex); i.Cmp(endIndex) <= 0; i.Add(i, big.NewInt(1)) {
		proposalInfoIndexed = append(proposalInfoIndexed, ProposalInfoIndexed{
			proposalInfos[index],
			int(i.Int64()),
		})
		index++
	}

	proposalInfosJson, err := json.Marshal(proposalInfoIndexed)

	if err != nil {
		return nil, fmt.Errorf("failed to Unmarshal Proposal Infos list: %v", err)
	}

	return proposalInfosJson, nil

}

func GetStateInfo(address string, rpc string, proposalId string) ([]byte, error) {

	conn, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Electroneum client: %v", err)
	}

	contract, err := NewStorage(common.HexToAddress(address), conn)

	if err != nil {
		return nil, fmt.Errorf("failed to instantiate Storage contract: %v", err)
	}

	bigInt, ok := new(big.Int).SetString(proposalId, 10)

	if !ok {
		return nil, fmt.Errorf("error converting string to big int")
	}

	info, err := contract.GetProposalInfo(nil, bigInt)

	if err != nil {
		return nil, fmt.Errorf("failed to GetProposalInfo contract: %v", err)
	}

	returnJson, err := json.Marshal(info)

	fmt.Printf(string(returnJson[:]))

	if err != nil {
		return nil, fmt.Errorf("cant unmarshal json %v", err)
	}

	return returnJson, nil
}
