package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type DialogRequestLeave struct{}

func (m DialogRequestLeave) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.DialogRequestLeave
}

func (m DialogRequestLeave) Serialized() (string, error) {
	return "", nil
}

func (m *DialogRequestLeave) Deserialize(extra string) error {
	return nil
}
