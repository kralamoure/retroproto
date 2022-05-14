package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type DialogRequestLeave struct{}

func (m DialogRequestLeave) MessageId() retroproto.MsgCliId {
	return retroproto.DialogRequestLeave
}

func (m DialogRequestLeave) MessageName() string {
	return "DialogRequestLeave"
}

func (m DialogRequestLeave) Serialized() (string, error) {
	return "", nil
}

func (m *DialogRequestLeave) Deserialize(extra string) error {
	return nil
}
