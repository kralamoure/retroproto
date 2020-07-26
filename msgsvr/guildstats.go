// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildStats struct{}

func (m GuildStats) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildStats
}

func (m GuildStats) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildStats) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
