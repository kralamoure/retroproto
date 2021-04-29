package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AksRPing struct {
	Value string
}

func (m AksRPing) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AksRPing
}

func (m AksRPing) Serialized() (string, error) {
	return m.Value, nil
}

func (m *AksRPing) Deserialize(extra string) error {
	m.Value = extra

	return nil
}
