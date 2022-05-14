// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildInfosTaxCollectorsMovement struct{}

func NewGuildInfosTaxCollectorsMovement(extra string) (GuildInfosTaxCollectorsMovement, error) {
	var m GuildInfosTaxCollectorsMovement

	if err := m.Deserialize(extra); err != nil {
		return GuildInfosTaxCollectorsMovement{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildInfosTaxCollectorsMovement) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildInfosTaxCollectorsMovement
}

func (m GuildInfosTaxCollectorsMovement) MessageName() string {
	return "GuildInfosTaxCollectorsMovement"
}

func (m GuildInfosTaxCollectorsMovement) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosTaxCollectorsMovement) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
