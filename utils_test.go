package identity_test

import (
	"math/big"
	"testing"

	identity "github.com/rarimo/rarime-mobile-identity-sdk"
	"github.com/stretchr/testify/assert"
)

func TestBytesArrayToBits(t *testing.T) {
	bytes := []byte{0x01, 0x02, 0x03}

	bits := identity.ByteArrayToBits(bytes)

	assert.Equal(t, []int64{0, 0, 0, 0, 0, 0, 0, 1}, bits[:8], "First byte")
	assert.Equal(t, []int64{0, 0, 0, 0, 0, 0, 1, 0}, bits[8:16], "Second byte")
	assert.Equal(t, []int64{0, 0, 0, 0, 0, 0, 1, 1}, bits[16:24], "Third byte")
}

func TestSmartChunking(t *testing.T) {
	bi, ok := new(big.Int).SetString("5f2aac3c296c6e61e078cc05c597d71fbab6406a8e0565f7d27e79141dfbf5e3ccec4275072fe6d949bad3ed409fd4bcdc9cd6d5115663e430e80bee938af27e5bcdec1729da2bf4a942f6d96ed4b58adf91bc2dbfee68d70e867c5b6f20038624e54becc0284720c826e0260e3152f23ecf313685b3f6801e840f261732106a23807a9d4aa56c481c69e86310587140aa5d6a7e291393eafeaeca7460ae2e7cc1c836868d9c7b12eeab75b09b042a7c6f735622fb0ca572c820f76e98094649f9ed1567225ad3beef330c16335ca78889de1f2db8183880fcd9a6df6ed25c718a79cef2fe26c33d86047a8e5fff781360fdc00798abc2663e12a9814da78839387e161d4a4795fa24835ac5f976faa5205c2c4dbde78a5245af89b394f0b49051aa261386655e2807d448a0986bba85c717a0ac5ec7957403febad91f5143c7d07101a2294d16a68ce93c55c1604076037526c6eda16921e85576abf5382eaadea0434f01eec6b51fcc15175f39f7651aa510566829f26e2dd940dd0a0204d3d4d4fdcf980f4b34ffb442955b23c72e1a5e1ece672a2939e90c29e1b03ed25e319926dd9462d0be47a453112d0c33a1226a91453482495bf540b49889a7c731d474961c4c651dc6edac39f083ac66e67ee8ff6851b5db92c86695d8df81d10f04de08ab5471af99640b80d102551d88a9b0305884d3b88f3ce0c708cbbcc45e", 16)
	if !ok {
		t.Error("Error parsing big int")
	}

	res := identity.SmartChunking(bi)

	t.Log(res)

	assert.Equal(t, 64, len(res), "Number of chunks")
}
