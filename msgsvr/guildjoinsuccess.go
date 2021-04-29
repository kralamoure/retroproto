// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildJoinSuccess struct{}

func (m GuildJoinSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildJoinSuccess
}

func (m GuildJoinSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildJoinSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
