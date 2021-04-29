// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildJoinSuccess struct{}

func (m GuildJoinSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildJoinSuccess
}

func (m GuildJoinSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildJoinSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
