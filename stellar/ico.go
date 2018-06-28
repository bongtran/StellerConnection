package stellar

import (
	"github.com/stellar/go/keypair"
)

func InitAccount() (error, keypair.Full) {
	pair, err := CreateKeyPair()

	if err != nil {
		return err, *pair
	}

	result := CreateAccount(pair.Address())

	return result, *pair
}

func ICO(issuingSeed string, distributionSeed string, code string, limit string) (string, error) {
	//issuingSeed := "SDE4DI3KZUNUH634NNQX7JGTMKBJHECHPHHVBLVYHRV2U6XPWVVVM6IV"
	//distributionSeed := "SATGL6FT2G2XW73VHEROXERRDXTXYVE7FJXO6S2KARXTY4HMETSF2JJA"

	return initAsset(issuingSeed, distributionSeed, code, limit)
}

func SendAsset(from string, to string, issuer, asset string, amount string) (string, error) {
	return SendAnAssetPaymentToThirdPerson(from, to, issuer, asset, amount)
}
