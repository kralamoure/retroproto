// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildGetInfosMembers struct{}

func (m GuildGetInfosMembers) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildGetInfosMembers
}

func (m GuildGetInfosMembers) Serialized() (string, error) {
	return "", nil
}

func (m *GuildGetInfosMembers) Deserialize(extra string) error {
	return nil
}
