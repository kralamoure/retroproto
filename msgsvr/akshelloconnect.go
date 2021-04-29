package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AksHelloConnect struct {
	Salt string
}

func (m AksHelloConnect) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AksHelloConnect
}

func (m AksHelloConnect) Serialized() (string, error) {
	return m.Salt, nil
}

func (m *AksHelloConnect) Deserialize(extra string) error {
	m.Salt = extra

	return nil
}
