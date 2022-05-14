package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AksQuickPong struct{}

func (m AksQuickPong) MessageId() retroproto.MsgSvrId {
	return retroproto.AksQuickPong
}

func (m AksQuickPong) MessageName() string {
	return "AksQuickPong"
}

func (m AksQuickPong) Serialized() (string, error) {
	return "", nil
}

func (m *AksQuickPong) Deserialize(extra string) error {
	return nil
}
