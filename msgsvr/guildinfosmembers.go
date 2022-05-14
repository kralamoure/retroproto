// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildInfosMembers struct{}

func (m GuildInfosMembers) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildInfosMembers
}

func (m GuildInfosMembers) MessageName() string {
	return "GuildInfosMembers"
}

func (m GuildInfosMembers) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosMembers) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
