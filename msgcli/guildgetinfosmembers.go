// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildGetInfosMembers struct{}

func (m GuildGetInfosMembers) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildGetInfosMembers
}

func (m GuildGetInfosMembers) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildGetInfosMembers) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
