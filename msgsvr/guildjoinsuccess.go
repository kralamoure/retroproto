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
	return "", nil
}

func (m *GuildJoinSuccess) Deserialize(extra string) error {
	return nil
}
