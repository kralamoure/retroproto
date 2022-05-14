package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AksRPing struct {
	Value string
}

func (m AksRPing) MessageId() retroproto.MsgSvrId {
	return retroproto.AksRPing
}

func (m AksRPing) MessageName() string {
	return "AksRPing"
}

func (m AksRPing) Serialized() (string, error) {
	return m.Value, nil
}

func (m *AksRPing) Deserialize(extra string) error {
	m.Value = extra

	return nil
}
