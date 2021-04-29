// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameExtraClip struct{}

func (m GameExtraClip) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameExtraClip
}

func (m GameExtraClip) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameExtraClip) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
