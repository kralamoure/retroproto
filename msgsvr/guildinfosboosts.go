// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildInfosBoosts struct{}

func (m GuildInfosBoosts) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GuildInfosBoosts
}

func (m GuildInfosBoosts) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosBoosts) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
