// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameFrameObjectExternal struct{}

func (m GameFrameObjectExternal) MessageId() retroproto.MsgSvrId {
	return retroproto.GameFrameObjectExternal
}

func (m GameFrameObjectExternal) MessageName() string {
	return "GameFrameObjectExternal"
}

func (m GameFrameObjectExternal) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFrameObjectExternal) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
