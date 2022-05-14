// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildLeaveTaxCollector struct{}

func NewGuildLeaveTaxCollector(extra string) (GuildLeaveTaxCollector, error) {
	var m GuildLeaveTaxCollector

	if err := m.Deserialize(extra); err != nil {
		return GuildLeaveTaxCollector{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildLeaveTaxCollector) MessageId() retroproto.MsgCliId {
	return retroproto.GuildLeaveTaxCollector
}

func (m GuildLeaveTaxCollector) MessageName() string {
	return "GuildLeaveTaxCollector"
}

func (m GuildLeaveTaxCollector) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildLeaveTaxCollector) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
