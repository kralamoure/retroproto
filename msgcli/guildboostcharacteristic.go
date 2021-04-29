// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildBoostCharacteristic struct{}

func (m GuildBoostCharacteristic) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildBoostCharacteristic
}

func (m GuildBoostCharacteristic) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildBoostCharacteristic) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
