package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type DialogRequestLeave struct{}

func (m DialogRequestLeave) ProtocolId() d1proto.MsgCliId {
	return d1proto.DialogRequestLeave
}

func (m DialogRequestLeave) Serialized() (string, error) {
	return "", nil
}

func (m *DialogRequestLeave) Deserialize(extra string) error {
	return nil
}
