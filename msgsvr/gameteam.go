// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameTeam struct{}

func (m GameTeam) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameTeam
}

func (m GameTeam) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameTeam) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
