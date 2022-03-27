package ed25519

import (
	"crypto/ed25519"
	"encoding/hex"
	"io"
)

type (
	PublicKey  ed25519.PublicKey
	PrivateKey ed25519.PrivateKey
)

func (k PublicKey) String() string {
	return hex.EncodeToString(k)
}

func (k PrivateKey) String() string {
	return hex.EncodeToString(k)
}

func Generate(rand io.Reader) (PublicKey, PrivateKey, error) {
	pub, priv, err := ed25519.GenerateKey(rand)
	if err != nil {
		return nil, nil, err
	}

	return PublicKey(pub), PrivateKey(priv), nil
}
