// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameFrameObjectExternal struct{}

func (m GameFrameObjectExternal) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameFrameObjectExternal
}

func (m GameFrameObjectExternal) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameFrameObjectExternal) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
