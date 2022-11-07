package enums

import "github.com/Amirilidan78/ethereum-wallet/util"

type ContractAddress string

func (ca ContractAddress) String() string {
	return string(ca)
}

func (ca ContractAddress) Bytes() []byte {
	bytes, _ := util.HexStringToBytes(string(ca))
	return bytes
}

func CreateContractAddress(contractAddressBase58 string) ContractAddress {
	return ContractAddress(contractAddressBase58)
}

const (
	GOERLI_USDT ContractAddress = "0x509Ee0d083DdF8AC028f2a56731412edD63223B9"
	GOERLI_USDC ContractAddress = "0x07865c6e87b9f70255377e024ace6630c1eaa37f"
)
