// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildCreate struct{}

func (m GuildCreate) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildCreate
}

func (m GuildCreate) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildCreate) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
