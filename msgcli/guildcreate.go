// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildCreate struct{}

func (m GuildCreate) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildCreate
}

func (m GuildCreate) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildCreate) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}