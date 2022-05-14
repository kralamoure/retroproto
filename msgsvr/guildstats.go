// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildStats struct{}

func (m GuildStats) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildStats
}

func (m GuildStats) MessageName() string {
	return "GuildStats"
}

func (m GuildStats) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildStats) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
