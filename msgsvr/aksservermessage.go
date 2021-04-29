package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AksServerMessage struct {
	// TODO
	Value string
}

func (m AksServerMessage) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AksServerMessage
}

func (m AksServerMessage) Serialized() (string, error) {
	return m.Value, nil
}

func (m *AksServerMessage) Deserialize(extra string) error {
	m.Value = extra

	return nil
}
