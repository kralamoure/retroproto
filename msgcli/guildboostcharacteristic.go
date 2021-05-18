// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildBoostCharacteristic struct{}

func (m GuildBoostCharacteristic) ProtocolId() retroproto.MsgCliId {
	return retroproto.GuildBoostCharacteristic
}

func (m GuildBoostCharacteristic) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildBoostCharacteristic) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
