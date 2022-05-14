package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AksServerMessage struct {
	// TODO
	Value string
}

func (m AksServerMessage) MessageId() retroproto.MsgSvrId {
	return retroproto.AksServerMessage
}

func (m AksServerMessage) MessageName() string {
	return "AksServerMessage"
}

func (m AksServerMessage) Serialized() (string, error) {
	return m.Value, nil
}

func (m *AksServerMessage) Deserialize(extra string) error {
	m.Value = extra

	return nil
}
