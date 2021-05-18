package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type DialogLeave struct{}

func (m DialogLeave) ProtocolId() retroproto.MsgSvrId {
	return retroproto.DialogLeave
}

func (m DialogLeave) Serialized() (string, error) {
	return "", nil
}

func (m *DialogLeave) Deserialize(extra string) error {
	return nil
}
