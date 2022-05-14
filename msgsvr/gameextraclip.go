// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameExtraClip struct{}

func (m GameExtraClip) MessageId() retroproto.MsgSvrId {
	return retroproto.GameExtraClip
}

func (m GameExtraClip) MessageName() string {
	return "GameExtraClip"
}

func (m GameExtraClip) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameExtraClip) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
