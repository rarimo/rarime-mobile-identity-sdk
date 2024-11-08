package brainpool

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/asn1"
	"encoding/pem"
	"math/big"
	"sync"
)

var (
	once                                                   sync.Once
	p160t1, p192t1, p224t1, p256t1, p320t1, p384t1, p512t1 *elliptic.CurveParams
	p160r1, p192r1, p224r1, p256r1, p320r1, p384r1, p512r1 *rcurve
)

func initAll() {
	initP160t1()
	initP192t1()
	initP224t1()
	initP256t1()
	initP320t1()
	initP384t1()
	initP512t1()

	initP160r1()
	initP192r1()
	initP224r1()
	initP256r1()
	initP320r1()
	initP384r1()
	initP512r1()
}

func initP160t1() {
	p160t1 = &elliptic.CurveParams{Name: "brainpoolP160t1"}
	p160t1.P, _ = new(big.Int).SetString("0xe95e4a5f737059dc60dfc7ad95b3d8139515620f", 0)
	p160t1.N, _ = new(big.Int).SetString("0xe95e4a5f737059dc60df5991d45029409e60fc09", 0)
	p160t1.B, _ = new(big.Int).SetString("0x7a556b6dae535b7b51ed2c4d7daa7a0b5c55f380", 0)
	p160t1.Gx, _ = new(big.Int).SetString("0xb199b13b9b34efc1397e64baeb05acc265ff2378", 0)
	p160t1.Gy, _ = new(big.Int).SetString("0xadd6718b7c7c1961f0991b842443772152c9e0ad", 0)
	p160t1.BitSize = 160
}

func initP192t1() {
	p192t1 = &elliptic.CurveParams{Name: "brainpoolP192t1"}
	p192t1.P, _ = new(big.Int).SetString("0xc302f41d932a36cda7a3463093d18db78fce476de1a86297", 0)
	p192t1.N, _ = new(big.Int).SetString("0xc302f41d932a36cda7a3462f9e9e916b5be8f1029ac4acc1", 0)
	p192t1.B, _ = new(big.Int).SetString("0x13d56ffaec78681e68f9deb43b35bec2fb68542e27897b79", 0)
	p192t1.Gx, _ = new(big.Int).SetString("0x3ae9e58c82f63c30282e1fe7bbf43fa72c446af6f4618129", 0)
	p192t1.Gy, _ = new(big.Int).SetString("0x97e2c5667c2223a902ab5ca449d0084b7e5b3de7ccc01c9", 0)
	p192t1.BitSize = 192
}

func initP224t1() {
	p224t1 = &elliptic.CurveParams{Name: "brainpoolP224t1"}
	p224t1.P, _ = new(big.Int).SetString("0xd7c134aa264366862a18302575d1d787b09f075797da89f57ec8c0ff", 0)
	p224t1.N, _ = new(big.Int).SetString("0xd7c134aa264366862a18302575d0fb98d116bc4b6ddebca3a5a7939f", 0)
	p224t1.B, _ = new(big.Int).SetString("0x4b337d934104cd7bef271bf60ced1ed20da14c08b3bb64f18a60888d", 0)
	p224t1.Gx, _ = new(big.Int).SetString("0x6ab1e344ce25ff3896424e7ffe14762ecb49f8928ac0c76029b4d580", 0)
	p224t1.Gy, _ = new(big.Int).SetString("0x374e9f5143e568cd23f3f4d7c0d4b1e41c8cc0d1c6abd5f1a46db4c", 0)
	p224t1.BitSize = 224
}

func initP256t1() {
	p256t1 = &elliptic.CurveParams{Name: "brainpoolP256t1"}
	p256t1.P, _ = new(big.Int).SetString("0xA9FB57DBA1EEA9BC3E660A909D838D726E3BF623D52620282013481D1F6E5377", 0)
	p256t1.N, _ = new(big.Int).SetString("0xA9FB57DBA1EEA9BC3E660A909D838D718C397AA3B561A6F7901E0E82974856A7", 0)
	p256t1.B, _ = new(big.Int).SetString("0x662C61C430D84EA4FE66A7733D0B76B7BF93EBC4AF2F49256AE58101FEE92B04", 0)
	p256t1.Gx, _ = new(big.Int).SetString("0xA3E8EB3CC1CFE7B7732213B23A656149AFA142C47AAFBC2B79A191562E1305F4", 0)
	p256t1.Gy, _ = new(big.Int).SetString("0x2D996C823439C56D7F7B22E14644417E69BCB6DE39D027001DABE8F35B25C9BE", 0)
	p256t1.BitSize = 256
}

func initP320t1() {
	p320t1 = &elliptic.CurveParams{Name: "brainpoolP320t1"}
	p320t1.P, _ = new(big.Int).SetString("0xd35e472036bc4fb7e13c785ed201e065f98fcfa6f6f40def4f92b9ec7893ec28fcd412b1f1b32e27", 0)
	p320t1.N, _ = new(big.Int).SetString("0xd35e472036bc4fb7e13c785ed201e065f98fcfa5b68f12a32d482ec7ee8658e98691555b44c59311", 0)
	p320t1.B, _ = new(big.Int).SetString("0xa7f561e038eb1ed560b3d147db782013064c19f27ed27c6780aaf77fb8a547ceb5b4fef422340353", 0)
	p320t1.Gx, _ = new(big.Int).SetString("0x925be9fb01afc6fb4d3e7d4990010f813408ab106c4f09cb7ee07868cc136fff3357f624a21bed52", 0)
	p320t1.Gy, _ = new(big.Int).SetString("0x63ba3a7a27483ebf6671dbef7abb30ebee084e58a0b077ad42a5a0989d1ee71b1b9bc0455fb0d2c3", 0)
	p320t1.BitSize = 320
}

func initP384t1() {
	p384t1 = &elliptic.CurveParams{Name: "brainpoolP384t1"}
	p384t1.P, _ = new(big.Int).SetString("0x8CB91E82A3386D280F5D6F7E50E641DF152F7109ED5456B412B1DA197FB71123ACD3A729901D1A71874700133107EC53", 0)
	p384t1.N, _ = new(big.Int).SetString("0x8CB91E82A3386D280F5D6F7E50E641DF152F7109ED5456B31F166E6CAC0425A7CF3AB6AF6B7FC3103B883202E9046565", 0)
	p384t1.B, _ = new(big.Int).SetString("0x7F519EADA7BDA81BD826DBA647910F8C4B9346ED8CCDC64E4B1ABD11756DCE1D2074AA263B88805CED70355A33B471EE", 0)
	p384t1.Gx, _ = new(big.Int).SetString("0x18DE98B02DB9A306F2AFCD7235F72A819B80AB12EBD653172476FECD462AABFFC4FF191B946A5F54D8D0AA2F418808CC", 0)
	p384t1.Gy, _ = new(big.Int).SetString("0x25AB056962D30651A114AFD2755AD336747F93475B7A1FCA3B88F2B6A208CCFE469408584DC2B2912675BF5B9E582928", 0)
	p384t1.BitSize = 384
}

func initP512t1() {
	p512t1 = &elliptic.CurveParams{Name: "brainpoolP512t1"}
	p512t1.P, _ = new(big.Int).SetString("0xAADD9DB8DBE9C48B3FD4E6AE33C9FC07CB308DB3B3C9D20ED6639CCA703308717D4D9B009BC66842AECDA12AE6A380E62881FF2F2D82C68528AA6056583A48F3", 0)
	p512t1.N, _ = new(big.Int).SetString("0xAADD9DB8DBE9C48B3FD4E6AE33C9FC07CB308DB3B3C9D20ED6639CCA70330870553E5C414CA92619418661197FAC10471DB1D381085DDADDB58796829CA90069", 0)
	p512t1.B, _ = new(big.Int).SetString("0x7CBBBCF9441CFAB76E1890E46884EAE321F70C0BCB4981527897504BEC3E36A62BCDFA2304976540F6450085F2DAE145C22553B465763689180EA2571867423E", 0)
	p512t1.Gx, _ = new(big.Int).SetString("0x640ECE5C12788717B9C1BA06CBC2A6FEBA85842458C56DDE9DB1758D39C0313D82BA51735CDB3EA499AA77A7D6943A64F7A3F25FE26F06B51BAA2696FA9035DA", 0)
	p512t1.Gy, _ = new(big.Int).SetString("0x5B534BD595F5AF0FA2C892376C84ACE1BB4E3019B71634C01131159CAE03CEE9D9932184BEEF216BD71DF2DADF86A627306ECFF96DBB8BACE198B61E00F8B332", 0)
	p512t1.BitSize = 512
}

func initP160r1() {
	twisted := p160t1
	params := &elliptic.CurveParams{
		Name:    "brainpoolP160r1",
		P:       twisted.P,
		N:       twisted.N,
		BitSize: twisted.BitSize,
	}
	params.B, _ = new(big.Int).SetString("0x1e589a8595423412134faa2dbdec95c8d8675e58", 0)
	params.Gx, _ = new(big.Int).SetString("0xbed5af16ea3f6a4f62938c4631eb5af7bdbcdbc3", 0)
	params.Gy, _ = new(big.Int).SetString("0x1667cb477a1a8ec338f94741669c976316da6321", 0)

	z, _ := new(big.Int).SetString("0x24DBFF5DEC9B986BBFE5295A29BFBAE45E0F5D0B", 0)
	p160r1 = newrcurve(twisted, params, z)
}

func initP192r1() {
	twisted := p192t1
	params := &elliptic.CurveParams{
		Name:    "brainpoolP192r1",
		P:       twisted.P,
		N:       twisted.N,
		BitSize: twisted.BitSize,
	}
	params.B, _ = new(big.Int).SetString("0x469a28ef7c28cca3dc721d044f4496bcca7ef4146fbf25c9", 0)
	params.Gx, _ = new(big.Int).SetString("0xc0a0647eaab6a48753b033c56cb0f0900a2f5c4853375fd6", 0)
	params.Gy, _ = new(big.Int).SetString("0x14b690866abd5bb88b5f4828c1490002e6773fa2fa299b8f", 0)
	z, _ := new(big.Int).SetString("0x1B6F5CC8DB4DC7AF19458A9CB80DC2295E5EB9C3732104CB", 0)
	p192r1 = newrcurve(twisted, params, z)
}

func initP224r1() {
	twisted := p224t1
	params := &elliptic.CurveParams{
		Name:    "brainpoolP224r1",
		P:       twisted.P,
		N:       twisted.N,
		BitSize: twisted.BitSize,
	}
	params.B, _ = new(big.Int).SetString("0x2580f63ccfe44138870713b1a92369e33e2135d266dbb372386c400b", 0)
	params.Gx, _ = new(big.Int).SetString("0xd9029ad2c7e5cf4340823b2a87dc68c9e4ce3174c1e6efdee12c07d", 0)
	params.Gy, _ = new(big.Int).SetString("0x58aa56f772c0726f24c6b89e4ecdac24354b9e99caa3f6d3761402cd", 0)
	z, _ := new(big.Int).SetString("0x2DF271E14427A346910CF7A2E6CFA7B3F484E5C2CCE1C8B730E28B3F", 0)
	p224r1 = newrcurve(twisted, params, z)
}

func initP256r1() {
	twisted := p256t1
	params := &elliptic.CurveParams{
		Name:    "brainpoolP256r1",
		P:       twisted.P,
		N:       twisted.N,
		BitSize: twisted.BitSize,
	}
	params.B, _ = new(big.Int).SetString("0x26dc5c6ce94a4b44f330b5d9bbd77cbf958416295cf7e1ce6bccdc18ff8c07b6", 0)
	params.Gx, _ = new(big.Int).SetString("0x8BD2AEB9CB7E57CB2C4B482FFC81B7AFB9DE27E1E3BD23C23A4453BD9ACE3262", 0)
	params.Gy, _ = new(big.Int).SetString("0x547EF835C3DAC4FD97F8461A14611DC9C27745132DED8E545C1D54C72F046997", 0)
	z, _ := new(big.Int).SetString("0x3E2D4BD9597B58639AE7AA669CAB9837CF5CF20A2C852D10F655668DFC150EF0", 0)
	p256r1 = newrcurve(twisted, params, z)
}

func initP320r1() {
	twisted := p320t1
	params := &elliptic.CurveParams{
		Name:    "brainpoolP320r1",
		P:       twisted.P,
		N:       twisted.N,
		BitSize: twisted.BitSize,
	}
	params.B, _ = new(big.Int).SetString("0x520883949dfdbc42d3ad198640688a6fe13f41349554b49acc31dccd884539816f5eb4ac8fb1f1a6", 0)
	params.Gx, _ = new(big.Int).SetString("0x43bd7e9afb53d8b85289bcc48ee5bfe6f20137d10a087eb6e7871e2a10a599c710af8d0d39e20611", 0)
	params.Gy, _ = new(big.Int).SetString("0x14fdd05545ec1cc8ab4093247f77275e0743ffed117182eaa9c77877aaac6ac7d35245d1692e8ee1", 0)
	z, _ := new(big.Int).SetString("0x15F75CAF668077F7E85B42EB01F0A81FF56ECD6191D55CB82B7D861458A18FEFC3E5AB7496F3C7B1", 0)
	p320r1 = newrcurve(twisted, params, z)
}

func initP384r1() {
	twisted := p384t1
	params := &elliptic.CurveParams{
		Name:    "brainpoolP384r1",
		P:       twisted.P,
		N:       twisted.N,
		BitSize: twisted.BitSize,
	}
	params.B, _ = new(big.Int).SetString("0x4a8c7dd22ce28268b39b55416f0447c2fb77de107dcd2a62e880ea53eeb62d57cb4390295dbc9943ab78696fa504c11", 0)
	params.Gx, _ = new(big.Int).SetString("0x1D1C64F068CF45FFA2A63A81B7C13F6B8847A3E77EF14FE3DB7FCAFE0CBD10E8E826E03436D646AAEF87B2E247D4AF1E", 0)
	params.Gy, _ = new(big.Int).SetString("0x8ABE1D7520F9C2A45CB1EB8E95CFD55262B70B29FEEC5864E19C054FF99129280E4646217791811142820341263C5315", 0)
	z, _ := new(big.Int).SetString("0x41DFE8DD399331F7166A66076734A89CD0D2BCDB7D068E44E1F378F41ECBAE97D2D63DBC87BCCDDCCC5DA39E8589291C", 0)
	p384r1 = newrcurve(twisted, params, z)
}

func initP512r1() {
	twisted := p512t1
	params := &elliptic.CurveParams{
		Name:    "brainpoolP512r1",
		P:       twisted.P,
		N:       twisted.N,
		BitSize: twisted.BitSize,
	}
	params.B, _ = new(big.Int).SetString("0x3df91610a83441caea9863bc2ded5d5aa8253aa10a2ef1c98b9ac8b57f1117a72bf2c7b9e7c1ac4d77fc94cadc083e67984050b75ebae5dd2809bd638016f723", 0)
	params.Gx, _ = new(big.Int).SetString("0x81AEE4BDD82ED9645A21322E9C4C6A9385ED9F70B5D916C1B43B62EEF4D0098EFF3B1F78E2D0D48D50D1687B93B97D5F7C6D5047406A5E688B352209BCB9F822", 0)
	params.Gy, _ = new(big.Int).SetString("0x7DDE385D566332ECC0EABFA9CF7822FDF209F70024A57B1AA000C55B881F8111B2DCDE494A5F485E5BCA4BD88A2763AED1CA2B2FA8F0540678CD1E0F3AD80892", 0)
	z, _ := new(big.Int).SetString("0x12EE58E6764838B69782136F0F2D3BA06E27695716054092E60A80BEDB212B64E585D90BCE13761F85C3F1D2A64E3BE8FEA2220F01EBA5EEB0F35DBD29D922AB", 0)
	p512r1 = newrcurve(twisted, params, z)
}

// P160t1 returns a Curve which implements the brainpoolP160t1 curve.
func P160t1() elliptic.Curve {
	once.Do(initAll)
	return p160t1
}

// P192t1 returns a Curve which implements the brainpoolP192t1 curve.
func P192t1() elliptic.Curve {
	once.Do(initAll)
	return p192t1
}

// P224t1 returns a Curve which implements the brainpoolP224t1 curve.
func P224t1() elliptic.Curve {
	once.Do(initAll)
	return p224t1
}

// P256t1 returns a Curve which implements the brainpoolP224t1 curve.
func P256t1() elliptic.Curve {
	once.Do(initAll)
	return p256t1
}

// P320t1 returns a Curve which implements the brainpoolP320t1 curve.
func P320t1() elliptic.Curve {
	once.Do(initAll)
	return p320t1
}

// P384t1 returns a Curve which implements the brainpoolP384t1 curve.
func P384t1() elliptic.Curve {
	once.Do(initAll)
	return p384t1
}

// P512t1 returns a Curve which implements the brainpoolP512t1 curve.
func P512t1() elliptic.Curve {
	once.Do(initAll)
	return p512t1
}

// P160r1 returns a Curve which implements the brainpoolP160r1 curve.
func P160r1() elliptic.Curve {
	once.Do(initAll)
	return p160r1
}

// P192r1 returns a Curve which implements the brainpoolP192r1 curve.
func P192r1() elliptic.Curve {
	once.Do(initAll)
	return p192r1
}

// P224r1 returns a Curve which implements the brainpoolP224r1 curve.
func P224r1() elliptic.Curve {
	once.Do(initAll)
	return p224r1
}

// P256r1 returns a Curve which implements the brainpoolP256r1 curve.
func P256r1() elliptic.Curve {
	once.Do(initAll)
	return p256r1
}

// P320r1 returns a Curve which implements the brainpoolP320r1 curve.
func P320r1() elliptic.Curve {
	once.Do(initAll)
	return p320r1
}

// P384r1 returns a Curve which implements the brainpoolP384r1 curve.
func P384r1() elliptic.Curve {
	once.Do(initAll)
	return p384r1
}

// P512r1 returns a Curve which implements the brainpoolP512r1 curve.
func P512r1() elliptic.Curve {
	once.Do(initAll)
	return p512r1
}

// GetPublicKeyFromPem extracts a public key from a PEM block.
func GetPublicKeyFromPem(pem *pem.Block) (*ecdsa.PublicKey, error) {
	var publicKeyInfo publicKeyInfo
	_, err := asn1.Unmarshal(pem.Bytes, &publicKeyInfo)
	if err != nil {
		return nil, err
	}

	b := new(big.Int).SetBytes(publicKeyInfo.Algorithm.Parameters.Curve.B)

	once.Do(initAll)

	var curve elliptic.Curve
	if b.Cmp(p160r1.params.B) == 0 {
		curve = P160r1()
	} else if b.Cmp(p192r1.params.B) == 0 {
		curve = P192t1()
	} else if b.Cmp(p224r1.params.B) == 0 {
		curve = P224r1()
	} else if b.Cmp(p256r1.params.B) == 0 {
		curve = P256r1()
	} else if b.Cmp(p320r1.params.B) == 0 {
		curve = P320r1()
	} else if b.Cmp(p384r1.params.B) == 0 {
		curve = P384r1()
	} else if b.Cmp(p512r1.params.B) == 0 {
		curve = P512r1()
	} else {
		return nil, nil
	}

	pubKeyData := publicKeyInfo.SubjectPublicKey.Bytes[1:]

	x := pubKeyData[0 : len(pubKeyData)/2]
	y := pubKeyData[len(pubKeyData)/2:]

	return &ecdsa.PublicKey{
		Curve: curve,
		X:     new(big.Int).SetBytes(x),
		Y:     new(big.Int).SetBytes(y),
	}, nil
}

// PublicKeyInfo represents the ASN.1 structure of a public key.
type publicKeyInfo struct {
	Algorithm        algorithmIdentifier
	SubjectPublicKey asn1.BitString
}

type algorithmIdentifier struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters ecParameters
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
	A           asn1.RawContent
	B           asn1.RawContent
}
