// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameTeam struct{}

func (m GameTeam) MessageId() retroproto.MsgSvrId {
	return retroproto.GameTeam
}

func (m GameTeam) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTeam) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
