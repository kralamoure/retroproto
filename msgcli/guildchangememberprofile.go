// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildChangeMemberProfile struct{}

func (m GuildChangeMemberProfile) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildChangeMemberProfile
}

func (m GuildChangeMemberProfile) Serialized() (string, error) {
	return "", nil
}

func (m *GuildChangeMemberProfile) Deserialize(extra string) error {
	return nil
}
