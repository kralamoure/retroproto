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
	return "", nil
}

func (m *GuildJoinError) Deserialize(extra string) error {
	return nil
}
