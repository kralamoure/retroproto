package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AksRPing struct {
	Value string
}

func (m AksRPing) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AksRPing
}

func (m AksRPing) Serialized() (string, error) {
	return m.Value, nil
}

func (m *AksRPing) Deserialize(extra string) error {
	m.Value = extra

	return nil
}
