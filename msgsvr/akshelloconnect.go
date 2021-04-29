package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AksHelloConnect struct {
	Salt string
}

func (m AksHelloConnect) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AksHelloConnect
}

func (m AksHelloConnect) Serialized() (string, error) {
	return m.Salt, nil
}

func (m *AksHelloConnect) Deserialize(extra string) error {
	m.Salt = extra

	return nil
}
