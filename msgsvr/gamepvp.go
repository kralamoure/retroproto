// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GamePVP struct{}

func (m GamePVP) MessageId() retroproto.MsgSvrId {
	return retroproto.GamePVP
}

func (m GamePVP) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GamePVP) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
