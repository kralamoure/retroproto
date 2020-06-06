package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AksHelloGame struct{}

func (m AksHelloGame) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AksHelloGame
}

func (m AksHelloGame) Serialized() (string, error) {
	return "", nil
}

func (m *AksHelloGame) Deserialize(extra string) error {
	return nil
}
