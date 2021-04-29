package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type DialogCreateError struct{}

func (m DialogCreateError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.DialogCreateError
}

func (m DialogCreateError) Serialized() (string, error) {
	return "", nil
}

func (m *DialogCreateError) Deserialize(extra string) error {
	return nil
}
