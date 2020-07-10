package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type DialogPause struct{}

func (m DialogPause) ProtocolId() d1proto.MsgSvrId {
	return d1proto.DialogPause
}

func (m DialogPause) Serialized() (string, error) {
	return "", nil
}

func (m *DialogPause) Deserialize(extra string) error {
	return nil
}
