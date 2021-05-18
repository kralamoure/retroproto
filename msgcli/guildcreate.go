// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildCreate struct{}

func (m GuildCreate) ProtocolId() retroproto.MsgCliId {
	return retroproto.GuildCreate
}

func (m GuildCreate) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildCreate) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
