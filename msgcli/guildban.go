// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildBan struct{}

func (m GuildBan) ProtocolId() retroproto.MsgCliId {
	return retroproto.GuildBan
}

func (m GuildBan) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildBan) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
