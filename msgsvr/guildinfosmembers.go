// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildInfosMembers struct{}

func (m GuildInfosMembers) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildInfosMembers
}

func (m GuildInfosMembers) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildInfosMembers) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
