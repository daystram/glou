package node

import (
	"io"

	"github.com/daystram/miu/crypto/ed25519"
)

type Identity struct {
	PublicAddress  ed25519.PublicKey
	PrivateAddress ed25519.PrivateKey
}

func NewIdentity(rand io.Reader) (*Identity, error) {
	pub, priv, err := ed25519.Generate(rand)
	if err != nil {
		return nil, err
	}

	return &Identity{
		PublicAddress:  pub,
		PrivateAddress: priv,
	}, nil
}
