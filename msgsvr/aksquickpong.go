package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AksQuickPong struct{}

func (m AksQuickPong) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AksQuickPong
}

func (m AksQuickPong) Serialized() (string, error) {
	return "", nil
}

func (m *AksQuickPong) Deserialize(extra string) error {
	return nil
}
