package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type DialogPause struct{}

func (m DialogPause) MessageId() retroproto.MsgSvrId {
	return retroproto.DialogPause
}

func (m DialogPause) MessageName() string {
	return "DialogPause"
}

func (m DialogPause) Serialized() (string, error) {
	return "", nil
}

func (m *DialogPause) Deserialize(extra string) error {
	return nil
}
