// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildJoinTaxCollector struct{}

func NewGuildJoinTaxCollector(extra string) (GuildJoinTaxCollector, error) {
	var m GuildJoinTaxCollector

	if err := m.Deserialize(extra); err != nil {
		return GuildJoinTaxCollector{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildJoinTaxCollector) MessageId() retroproto.MsgCliId {
	return retroproto.GuildJoinTaxCollector
}

func (m GuildJoinTaxCollector) MessageName() string {
	return "GuildJoinTaxCollector"
}

func (m GuildJoinTaxCollector) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildJoinTaxCollector) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
