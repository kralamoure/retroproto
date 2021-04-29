package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AksQuickPong struct{}

func (m AksQuickPong) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AksQuickPong
}

func (m AksQuickPong) Serialized() (string, error) {
	return "", nil
}

func (m *AksQuickPong) Deserialize(extra string) error {
	return nil
}
