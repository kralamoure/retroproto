package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type DialogPause struct{}

func (m DialogPause) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.DialogPause
}

func (m DialogPause) Serialized() (string, error) {
	return "", nil
}

func (m *DialogPause) Deserialize(extra string) error {
	return nil
}
