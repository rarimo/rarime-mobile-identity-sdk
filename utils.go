package identity

import (
	"encoding/hex"
	"math/big"

	"github.com/iden3/go-iden3-crypto/babyjub"
)

// NewBJJSecretKey generates a new secret key for the Baby JubJub curve.
func NewBJJSecretKey() string {
	secretKey := babyjub.NewRandPrivKey()

	return hex.EncodeToString(secretKey[:])
}

// ByteArrayToBits converts a byte array to a bit array.
func ByteArrayToBits(bytes []byte) []int64 {
	var bits []int64
	for _, b := range bytes {
		for i := 0; i < 8; i++ {
			bits = append(bits, int64(b>>uint(7-i)&1))
		}
	}

	return bits
}

// SmartChunking splits a big.Int into chunks of 8 bytes.
//
// It does its best to split the big.Int into chunks of 8 bytes,
// but it may not be perfect though it smart (I heard... do not believe everything you hear).
func SmartChunking(x *big.Int) []string {
	var res []string

	mod := big.NewInt(1)
	for i := 0; i < 64; i++ {
		mod.Mul(mod, big.NewInt(2))
	}
	for i := 0; i < 64; i++ {
		chunk := new(big.Int).Mod(x, mod)

		res = append(res, chunk.String())
		x.Div(x, mod)
	}

	return res
}

// def bigint_to_array(n: int, k: int, x: str) -> list[str]:
//     # Initialize mod to 1 (Python's int can handle arbitrarily large numbers)
//     mod = 1
//     for _ in range(n):
//         mod *= 2

//     # Initialize the return list
//     ret = []
//     x_temp = x
//     for _ in range(k):
//         # Append x_temp mod mod to the list
//         ret.append(str(x_temp % mod))
//         # Divide x_temp by mod for the next iteration
//         x_temp //= mod  # Use integer division in Python

//     return ret
