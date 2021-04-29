// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildJoinDistantSuccess struct{}

func (m GuildJoinDistantSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildJoinDistantSuccess
}

func (m GuildJoinDistantSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildJoinDistantSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
