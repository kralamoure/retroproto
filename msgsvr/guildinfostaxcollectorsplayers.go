// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildInfosTaxCollectorsPlayers struct{}

func NewGuildInfosTaxCollectorsPlayers(extra string) (GuildInfosTaxCollectorsPlayers, error) {
	var m GuildInfosTaxCollectorsPlayers

	if err := m.Deserialize(extra); err != nil {
		return GuildInfosTaxCollectorsPlayers{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildInfosTaxCollectorsPlayers) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildInfosTaxCollectorsPlayers
}

func (m GuildInfosTaxCollectorsPlayers) MessageName() string {
	return "GuildInfosTaxCollectorsPlayers"
}

func (m GuildInfosTaxCollectorsPlayers) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosTaxCollectorsPlayers) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
