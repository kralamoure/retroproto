package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type DialogCreateError struct{}

func (m DialogCreateError) MessageId() retroproto.MsgSvrId {
	return retroproto.DialogCreateError
}

func (m DialogCreateError) MessageName() string {
	return "DialogCreateError"
}

func (m DialogCreateError) Serialized() (string, error) {
	return "", nil
}

func (m *DialogCreateError) Deserialize(extra string) error {
	return nil
}
