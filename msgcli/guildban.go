// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildBan struct{}

func (m GuildBan) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildBan
}

func (m GuildBan) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildBan) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
