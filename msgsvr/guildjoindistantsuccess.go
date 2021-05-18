// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildJoinDistantSuccess struct{}

func (m GuildJoinDistantSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GuildJoinDistantSuccess
}

func (m GuildJoinDistantSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildJoinDistantSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
