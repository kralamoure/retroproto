package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type DialogCreateError struct{}

func (m DialogCreateError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.DialogCreateError
}

func (m DialogCreateError) Serialized() (string, error) {
	return "", nil
}

func (m *DialogCreateError) Deserialize(extra string) error {
	return nil
}
