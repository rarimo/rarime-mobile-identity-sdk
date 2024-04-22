package identity

// SMTProof represents a sparse merkle tree proof.
type SMTProof struct {
	Root     []byte   `json:"root"`
	Siblings [][]byte `json:"siblings"`
}
