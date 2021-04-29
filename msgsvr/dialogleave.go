package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type DialogLeave struct{}

func (m DialogLeave) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.DialogLeave
}

func (m DialogLeave) Serialized() (string, error) {
	return "", nil
}

func (m *DialogLeave) Deserialize(extra string) error {
	return nil
}
