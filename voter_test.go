package identity_test

import (
	"encoding/json"
	identity "github.com/rarimo/rarime-mobile-identity-sdk"
	"testing"
)

func TestVoter(t *testing.T) {
	print("TestVoter")
	rpc_url := "https://rpc.qtestnet.org"
	proposalsStateContractAddress := "0xf4B99A3891D0a64A0bc3bB8642242E6A01e104e2"

	proposalInfoBytes, err := identity.GetProposalInfo(proposalsStateContractAddress, rpc_url, "17")
	if err != nil {
		t.Errorf("GetProposalInfo() failed: %v", err)
	}

	proposalInfo := new(identity.ProposalInfo)
	if err := json.Unmarshal(proposalInfoBytes, proposalInfo); err != nil {
		t.Errorf("proto.Unmarshal() failed: %v", err)
	}

	t.Logf("proposalInfo: %v\n", proposalInfo)
}
