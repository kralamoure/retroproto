// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildGetInfosMembers struct{}

func (m GuildGetInfosMembers) MessageId() retroproto.MsgCliId {
	return retroproto.GuildGetInfosMembers
}

func (m GuildGetInfosMembers) MessageName() string {
	return "GuildGetInfosMembers"
}

func (m GuildGetInfosMembers) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildGetInfosMembers) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
