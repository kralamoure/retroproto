// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildBoostCharacteristic struct{}

func (m GuildBoostCharacteristic) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildBoostCharacteristic
}

func (m GuildBoostCharacteristic) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildBoostCharacteristic) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
