// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildChangeMemberProfile struct{}

func (m GuildChangeMemberProfile) MessageId() retroproto.MsgCliId {
	return retroproto.GuildChangeMemberProfile
}

func (m GuildChangeMemberProfile) MessageName() string {
	return "GuildChangeMemberProfile"
}

func (m GuildChangeMemberProfile) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildChangeMemberProfile) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
