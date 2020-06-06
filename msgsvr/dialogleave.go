package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type DialogLeave struct{}

func (m DialogLeave) ProtocolId() d1proto.MsgSvrId {
	return d1proto.DialogLeave
}

func (m DialogLeave) Serialized() (string, error) {
	return "", nil
}

func (m *DialogLeave) Deserialize(extra string) error {
	return nil
}
