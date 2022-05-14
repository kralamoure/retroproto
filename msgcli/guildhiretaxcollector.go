// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildHireTaxCollector struct{}

func NewGuildHireTaxCollector(extra string) (GuildHireTaxCollector, error) {
	var m GuildHireTaxCollector

	if err := m.Deserialize(extra); err != nil {
		return GuildHireTaxCollector{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildHireTaxCollector) MessageId() retroproto.MsgCliId {
	return retroproto.GuildHireTaxCollector
}

func (m GuildHireTaxCollector) MessageName() string {
	return "GuildHireTaxCollector"
}

func (m GuildHireTaxCollector) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildHireTaxCollector) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
