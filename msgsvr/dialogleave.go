package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type DialogLeave struct{}

func (m DialogLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.DialogLeave
}

func (m DialogLeave) MessageName() string {
	return "DialogLeave"
}

func (m DialogLeave) Serialized() (string, error) {
	return "", nil
}

func (m *DialogLeave) Deserialize(extra string) error {
	return nil
}
