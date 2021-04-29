package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AksServerWillDisconnect struct{}

func (m AksServerWillDisconnect) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AksServerWillDisconnect
}

func (m AksServerWillDisconnect) Serialized() (string, error) {
	return "", nil
}

func (m *AksServerWillDisconnect) Deserialize(extra string) error {
	return nil
}
