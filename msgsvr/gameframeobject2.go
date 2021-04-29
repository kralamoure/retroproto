// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameFrameObject2 struct{}

func (m GameFrameObject2) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameFrameObject2
}

func (m GameFrameObject2) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameFrameObject2) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
