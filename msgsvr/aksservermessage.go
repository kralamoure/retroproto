package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AksServerMessage struct {
	// TODO
	Value string
}

func (m AksServerMessage) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AksServerMessage
}

func (m AksServerMessage) Serialized() (string, error) {
	return m.Value, nil
}

func (m *AksServerMessage) Deserialize(extra string) error {
	m.Value = extra

	return nil
}
