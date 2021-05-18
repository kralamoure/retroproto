package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AksServerWillDisconnect struct{}

func (m AksServerWillDisconnect) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AksServerWillDisconnect
}

func (m AksServerWillDisconnect) Serialized() (string, error) {
	return "", nil
}

func (m *AksServerWillDisconnect) Deserialize(extra string) error {
	return nil
}
