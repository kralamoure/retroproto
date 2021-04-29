// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildBoostSpell struct{}

func (m GuildBoostSpell) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildBoostSpell
}

func (m GuildBoostSpell) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildBoostSpell) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
