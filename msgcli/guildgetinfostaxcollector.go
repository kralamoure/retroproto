// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildGetInfosTaxCollector struct{}

func NewGuildGetInfosTaxCollector(extra string) (GuildGetInfosTaxCollector, error) {
	var m GuildGetInfosTaxCollector

	if err := m.Deserialize(extra); err != nil {
		return GuildGetInfosTaxCollector{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildGetInfosTaxCollector) MessageId() retroproto.MsgCliId {
	return retroproto.GuildGetInfosTaxCollector
}

func (m GuildGetInfosTaxCollector) MessageName() string {
	return "GuildGetInfosTaxCollector"
}

func (m GuildGetInfosTaxCollector) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildGetInfosTaxCollector) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
