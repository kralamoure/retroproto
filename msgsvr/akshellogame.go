package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AksHelloGame struct{}

func (m AksHelloGame) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AksHelloGame
}

func (m AksHelloGame) Serialized() (string, error) {
	return "", nil
}

func (m *AksHelloGame) Deserialize(extra string) error {
	return nil
}
