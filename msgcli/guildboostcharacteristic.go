// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildBoostCharacteristic struct{}

func NewGuildBoostCharacteristic(extra string) (GuildBoostCharacteristic, error) {
	var m GuildBoostCharacteristic

	if err := m.Deserialize(extra); err != nil {
		return GuildBoostCharacteristic{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildBoostCharacteristic) MessageId() retroproto.MsgCliId {
	return retroproto.GuildBoostCharacteristic
}

func (m GuildBoostCharacteristic) MessageName() string {
	return "GuildBoostCharacteristic"
}

func (m GuildBoostCharacteristic) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildBoostCharacteristic) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
