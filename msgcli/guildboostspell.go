// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildBoostSpell struct{}

func (m GuildBoostSpell) ProtocolId() retroproto.MsgCliId {
	return retroproto.GuildBoostSpell
}

func (m GuildBoostSpell) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildBoostSpell) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
