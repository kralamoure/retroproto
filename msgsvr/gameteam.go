// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameTeam struct{}

func (m GameTeam) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameTeam
}

func (m GameTeam) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameTeam) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
