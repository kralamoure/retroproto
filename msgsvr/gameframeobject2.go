// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameFrameObject2 struct{}

func (m GameFrameObject2) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GameFrameObject2
}

func (m GameFrameObject2) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFrameObject2) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
