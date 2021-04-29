package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AksHelloGame struct{}

func (m AksHelloGame) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AksHelloGame
}

func (m AksHelloGame) Serialized() (string, error) {
	return "", nil
}

func (m *AksHelloGame) Deserialize(extra string) error {
	return nil
}
