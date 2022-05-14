// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildInfosTaxCollectorsAttackers struct{}

func NewGuildInfosTaxCollectorsAttackers(extra string) (GuildInfosTaxCollectorsAttackers, error) {
	var m GuildInfosTaxCollectorsAttackers

	if err := m.Deserialize(extra); err != nil {
		return GuildInfosTaxCollectorsAttackers{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildInfosTaxCollectorsAttackers) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildInfosTaxCollectorsAttackers
}

func (m GuildInfosTaxCollectorsAttackers) MessageName() string {
	return "GuildInfosTaxCollectorsAttackers"
}

func (m GuildInfosTaxCollectorsAttackers) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosTaxCollectorsAttackers) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
