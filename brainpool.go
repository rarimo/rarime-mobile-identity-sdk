package identity

import (
	"crypto/elliptic"
	"encoding/asn1"
	"math/big"
)

var brainpoolP160r1CurveParams = getCurveParamsBrainpoolP160r1()
var brainpoolP192r1CurveParams = getCurveParamsBrainpoolP192r1()
var brainpoolP224r1CurveParams = getCurveParamsBrainpoolP224r1()
var brainpoolP256r1CurveParams = getCurveParamsBrainpoolP256r1()
var brainpoolP320r1CurveParams = getCurveParamsBrainpoolP320r1()
var brainpoolP384r1CurveParams = getCurveParamsBrainpoolP384r1()
var brainpoolP512r1CurveParams = getCurveParamsBrainpoolP512r1()

func getCurveParamsBrainpoolP160r1() *elliptic.CurveParams {
	p, _ := new(big.Int).SetString("0xe95e4a5f737059dc60dfc7ad95b3d8139515620f", 0)
	n, _ := new(big.Int).SetString("0xe95e4a5f737059dc60df5991d45029409e60fc09", 0)
	b, _ := new(big.Int).SetString("0x1e589a8595423412134faa2dbdec95c8d8675e58", 0)
	gX, _ := new(big.Int).SetString("0xbed5af16ea3f6a4f62938c4631eb5af7bdbcdbc3", 0)
	gY, _ := new(big.Int).SetString("0x1667cb477a1a8ec338f94741669c976316da6321", 0)

	return &elliptic.CurveParams{
		P:       p,
		N:       n,
		B:       b,
		Gx:      gX,
		Gy:      gY,
		BitSize: 160,
		Name:    "brainpoolP160r1",
	}
}

func getCurveParamsBrainpoolP192r1() *elliptic.CurveParams {
	p, _ := new(big.Int).SetString("0xc302f41d932a36cda7a3463093d18db78fce476de1a86297", 0)
	n, _ := new(big.Int).SetString("0xc302f41d932a36cda7a3462f9e9e916b5be8f1029ac4acc1", 0)
	b, _ := new(big.Int).SetString("0x469a28ef7c28cca3dc721d044f4496bcca7ef4146fbf25c9", 0)
	gX, _ := new(big.Int).SetString("0xc0a0647eaab6a48753b033c56cb0f0900a2f5c4853375fd6", 0)
	gY, _ := new(big.Int).SetString("0x14b690866abd5bb88b5f4828c1490002e6773fa2fa299b8f", 0)

	return &elliptic.CurveParams{
		P:       p,
		N:       n,
		B:       b,
		Gx:      gX,
		Gy:      gY,
		BitSize: 192,
		Name:    "brainpoolP192r1",
	}
}

func getCurveParamsBrainpoolP224r1() *elliptic.CurveParams {
	p, _ := new(big.Int).SetString("0xd7c134aa264366862a18302575d1d787b09f075797da89f57ec8c0ff", 0)
	n, _ := new(big.Int).SetString("0xd7c134aa264366862a18302575d0fb98d116bc4b6ddebca3a5a7939f", 0)
	b, _ := new(big.Int).SetString("0x2580f63ccfe44138870713b1a92369e33e2135d266dbb372386c400b", 0)
	gX, _ := new(big.Int).SetString("0xd9029ad2c7e5cf4340823b2a87dc68c9e4ce3174c1e6efdee12c07d", 0)
	gY, _ := new(big.Int).SetString("0x58aa56f772c0726f24c6b89e4ecdac24354b9e99caa3f6d3761402cd", 0)

	return &elliptic.CurveParams{
		P:       p,
		N:       n,
		B:       b,
		Gx:      gX,
		Gy:      gY,
		BitSize: 224,
		Name:    "brainpoolP224r1",
	}
}

func getCurveParamsBrainpoolP256r1() *elliptic.CurveParams {
	p, _ := new(big.Int).SetString("0xa9fb57dba1eea9bc3e660a909d838d726e3bf623d52620282013481d1f6e5377", 0)
	n, _ := new(big.Int).SetString("0xa9fb57dba1eea9bc3e660a909d838d718c397aa3b561a6f7901e0e82974856a7", 0)
	b, _ := new(big.Int).SetString("0x26dc5c6ce94a4b44f330b5d9bbd77cbf958416295cf7e1ce6bccdc18ff8c07b6", 0)
	gX, _ := new(big.Int).SetString("0x8bd2aeb9cb7e57cb2c4b482ffc81b7afb9de27e1e3bd23c23a4453bd9ace3262", 0)
	gY, _ := new(big.Int).SetString("0x547ef835c3dac4fd97f8461a14611dc9c27745132ded8e545c1d54c72f046997", 0)

	return &elliptic.CurveParams{
		P:       p,
		N:       n,
		B:       b,
		Gx:      gX,
		Gy:      gY,
		BitSize: 256,
		Name:    "brainpoolP256r1",
	}
}

func getCurveParamsBrainpoolP320r1() *elliptic.CurveParams {
	p, _ := new(big.Int).SetString("0xd35e472036bc4fb7e13c785ed201e065f98fcfa6f6f40def4f92b9ec7893ec28fcd412b1f1b32e27", 0)
	n, _ := new(big.Int).SetString("0xd35e472036bc4fb7e13c785ed201e065f98fcfa5b68f12a32d482ec7ee8658e98691555b44c59311", 0)
	b, _ := new(big.Int).SetString("0x520883949dfdbc42d3ad198640688a6fe13f41349554b49acc31dccd884539816f5eb4ac8fb1f1a6", 0)
	gX, _ := new(big.Int).SetString("0x43bd7e9afb53d8b85289bcc48ee5bfe6f20137d10a087eb6e7871e2a10a599c710af8d0d39e20611", 0)
	gY, _ := new(big.Int).SetString("0x14fdd05545ec1cc8ab4093247f77275e0743ffed117182eaa9c77877aaac6ac7d35245d1692e8ee1", 0)

	return &elliptic.CurveParams{
		P:       p,
		N:       n,
		B:       b,
		Gx:      gX,
		Gy:      gY,
		BitSize: 320,
		Name:    "brainpoolP320r1",
	}
}

func getCurveParamsBrainpoolP384r1() *elliptic.CurveParams {
	p, _ := new(big.Int).SetString("0x8cb91e82a3386d280f5d6f7e50e641df152f7109ed5456b412b1da197fb71123acd3a729901d1a71874700133107ec53", 0)
	n, _ := new(big.Int).SetString("0x8cb91e82a3386d280f5d6f7e50e641df152f7109ed5456b31f166e6cac0425a7cf3ab6af6b7fc3103b883202e9046565", 0)
	b, _ := new(big.Int).SetString("0x4a8c7dd22ce28268b39b55416f0447c2fb77de107dcd2a62e880ea53eeb62d57cb4390295dbc9943ab78696fa504c11", 0)
	gX, _ := new(big.Int).SetString("0x1d1c64f068cf45ffa2a63a81b7c13f6b8847a3e77ef14fe3db7fcafe0cbd10e8e826e03436d646aaef87b2e247d4af1e", 0)
	gY, _ := new(big.Int).SetString("0x8abe1d7520f9c2a45cb1eb8e95cfd55262b70b29feec5864e19c054ff99129280e4646217791811142820341263c5315", 0)

	return &elliptic.CurveParams{
		P:       p,
		N:       n,
		B:       b,
		Gx:      gX,
		Gy:      gY,
		BitSize: 384,
		Name:    "brainpoolP384r1",
	}
}

func getCurveParamsBrainpoolP512r1() *elliptic.CurveParams {
	p, _ := new(big.Int).SetString("0xaadd9db8dbe9c48b3fd4e6ae33c9fc07cb308db3b3c9d20ed6639cca703308717d4d9b009bc66842aecda12ae6a380e62881ff2f2d82c68528aa6056583a48f3", 0)
	n, _ := new(big.Int).SetString("0xaadd9db8dbe9c48b3fd4e6ae33c9fc07cb308db3b3c9d20ed6639cca70330870553e5c414ca92619418661197fac10471db1d381085ddaddb58796829ca90069", 0)
	b, _ := new(big.Int).SetString("0x3df91610a83441caea9863bc2ded5d5aa8253aa10a2ef1c98b9ac8b57f1117a72bf2c7b9e7c1ac4d77fc94cadc083e67984050b75ebae5dd2809bd638016f723", 0)
	gX, _ := new(big.Int).SetString("0x81aee4bdd82ed9645a21322e9c4c6a9385ed9f70b5d916c1b43b62eef4d0098eff3b1f78e2d0d48d50d1687b93b97d5f7c6d5047406a5e688b352209bcb9f822", 0)
	gY, _ := new(big.Int).SetString("0x7dde385d566332ecc0eabfa9cf7822fdf209f70024a57b1aa000c55b881f8111b2dcde494a5f485e5bca4bd88a2763aed1ca2b2fa8f0540678cd1e0f3ad80892", 0)

	return &elliptic.CurveParams{
		P:       p,
		N:       n,
		B:       b,
		Gx:      gX,
		Gy:      gY,
		BitSize: 512,
		Name:    "brainpoolP512r1",
	}
}

type algorithmIdentifier struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters ecParameters
}

type publicKeyInfo struct {
	Algorithm        algorithmIdentifier
	SubjectPublicKey asn1.BitString
}

type ecParameters struct {
	Version *big.Int
	FieldID fieldID
	Curve   curve
	G       asn1.RawContent
	N       *big.Int
	H       *big.Int
}

type fieldID struct {
	FieldType asn1.ObjectIdentifier
	Data      *big.Int
}

type curve struct {
	Placeholder asn1.RawContent
	X           asn1.RawContent
	Y           asn1.RawContent
}
