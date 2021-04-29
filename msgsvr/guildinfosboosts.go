// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildInfosBoosts struct{}

func (m GuildInfosBoosts) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildInfosBoosts
}

func (m GuildInfosBoosts) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildInfosBoosts) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
