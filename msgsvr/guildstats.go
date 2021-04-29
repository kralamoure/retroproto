// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildStats struct{}

func (m GuildStats) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildStats
}

func (m GuildStats) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildStats) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
