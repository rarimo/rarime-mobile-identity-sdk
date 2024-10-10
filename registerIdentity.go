package identity

import "encoding/json"

// RegisterIdentityInputs represents the inputs for the registerIdentity circuit.
type RegisterIdentityInputs struct {
	SkIdentity                   string   `json:"skIdentity"`
	EncapsulatedContent          []int64  `json:"encapsulatedContent"`
	SignedAttributes             []int64  `json:"signedAttributes"`
	Signature                    []string `json:"signature"`
	Pubkey                       []string `json:"pubkey"`
	Dg1                          []int64  `json:"dg1"`
	Dg15                         []int64  `json:"dg15"`
	SlaveMerleRoot               string   `json:"slaveMerkleRoot"`
	SlaveMerkleInclusionBranches []string `json:"slaveMerkleInclusionBranches"`
}

// Marshal returns the JSON representation of the inputs.
func (r *RegisterIdentityInputs) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// PassportCertificateInputs represents the inputs for the passport certificate circuit.
type PassportCertificateInputs struct {
	SlaveSignedAttributes []int64  `json:"slaveSignedAttributes"`
	SlaveSignature        []string `json:"slaveSignature"`
	MasterModulus         []string `json:"masterModulus"`
}
