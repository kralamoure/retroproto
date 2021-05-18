package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AksHelloConnect struct {
	Salt string
}

func (m AksHelloConnect) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AksHelloConnect
}

func (m AksHelloConnect) Serialized() (string, error) {
	return m.Salt, nil
}

func (m *AksHelloConnect) Deserialize(extra string) error {
	m.Salt = extra

	return nil
}
