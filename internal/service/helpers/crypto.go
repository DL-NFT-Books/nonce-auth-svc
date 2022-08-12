package helpers

import (
	"strings"

	"gitlab.com/tokene/nonce-auth-svc/internal/service/util"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
)

func NonceToHash(nonce *data.Nonce) []byte {
	message := util.PrefixNonceMessage(nonce.Message)
	hash := crypto.Keccak256Hash([]byte(message)).Bytes()
	return hash
}

func DecodeSignature(signature string) ([]byte, error) {
	signatureBytes, err := hexutil.Decode(signature)
	if err != nil {
		return nil, err
	}
	if len(signatureBytes) != 65 {
		return nil, errors.New("bad signature length")
	}
	if signatureBytes[64] == 0 || signatureBytes[64] == 1 {
		signatureBytes[64] = signatureBytes[64] + 27
	}
	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	if signatureBytes[64] != 27 && signatureBytes[64] != 28 {
		return nil, errors.New("bad recovery byte")
	}
	signatureBytes[64] -= 27

	return signatureBytes, nil
}

func VerifySignature(hash []byte, signature string, addresses ...string) error {
	if len(addresses) < 1 {
		return errors.New("no addresses provided for signature verification")
	}
	signatureBytes, err := DecodeSignature(signature)
	if err != nil {
		return errors.Wrap(err, "failed to decode signature")
	}

	// FIXME(Yaroslav Panasenko): Are you sure that verify signature (crypto.VerifySignature) is not required?

	recoveredPubkey, err := crypto.SigToPub(hash, signatureBytes)
	if err != nil {
		return errors.Wrap(err, "failed to recover pubkey from signed message")
	}

	matched := false
	for _, address := range addresses {
		if strings.ToLower(address) == strings.ToLower(crypto.PubkeyToAddress(*recoveredPubkey).Hex()) {
			matched = true
		}
	}

	if !matched {
		return errors.New("recovered address didn't match any of the given ones")
	}

	return nil
}
