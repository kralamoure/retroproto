// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildLeaveTaxInterface struct{}

func NewGuildLeaveTaxInterface(extra string) (GuildLeaveTaxInterface, error) {
	var m GuildLeaveTaxInterface

	if err := m.Deserialize(extra); err != nil {
		return GuildLeaveTaxInterface{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildLeaveTaxInterface) MessageId() retroproto.MsgCliId {
	return retroproto.GuildLeaveTaxInterface
}

func (m GuildLeaveTaxInterface) MessageName() string {
	return "GuildLeaveTaxInterface"
}

func (m GuildLeaveTaxInterface) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildLeaveTaxInterface) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
