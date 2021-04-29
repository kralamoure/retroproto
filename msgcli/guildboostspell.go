// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildBoostSpell struct{}

func (m GuildBoostSpell) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildBoostSpell
}

func (m GuildBoostSpell) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildBoostSpell) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
