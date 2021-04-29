package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AksServerWillDisconnect struct{}

func (m AksServerWillDisconnect) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AksServerWillDisconnect
}

func (m AksServerWillDisconnect) Serialized() (string, error) {
	return "", nil
}

func (m *AksServerWillDisconnect) Deserialize(extra string) error {
	return nil
}
