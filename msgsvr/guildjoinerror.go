// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildJoinError struct{}

func (m GuildJoinError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildJoinError
}

func (m GuildJoinError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildJoinError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
