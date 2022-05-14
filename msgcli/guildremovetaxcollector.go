// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildRemoveTaxCollector struct{}

func NewGuildRemoveTaxCollector(extra string) (GuildRemoveTaxCollector, error) {
	var m GuildRemoveTaxCollector

	if err := m.Deserialize(extra); err != nil {
		return GuildRemoveTaxCollector{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildRemoveTaxCollector) MessageId() retroproto.MsgCliId {
	return retroproto.GuildRemoveTaxCollector
}

func (m GuildRemoveTaxCollector) MessageName() string {
	return "GuildRemoveTaxCollector"
}

func (m GuildRemoveTaxCollector) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildRemoveTaxCollector) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
