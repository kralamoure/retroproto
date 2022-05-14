// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildGetInfosGeneral struct{}

func NewGuildGetInfosGeneral(extra string) (GuildGetInfosGeneral, error) {
	var m GuildGetInfosGeneral

	if err := m.Deserialize(extra); err != nil {
		return GuildGetInfosGeneral{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildGetInfosGeneral) MessageId() retroproto.MsgCliId {
	return retroproto.GuildGetInfosGeneral
}

func (m GuildGetInfosGeneral) MessageName() string {
	return "GuildGetInfosGeneral"
}

func (m GuildGetInfosGeneral) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildGetInfosGeneral) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
