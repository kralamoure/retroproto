// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildJoinError struct{}

func (m GuildJoinError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GuildJoinError
}

func (m GuildJoinError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildJoinError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
