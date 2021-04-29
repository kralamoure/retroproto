// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildBan struct{}

func (m GuildBan) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildBan
}

func (m GuildBan) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildBan) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
