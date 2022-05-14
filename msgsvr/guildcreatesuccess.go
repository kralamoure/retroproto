// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildCreateSuccess struct{}

func NewGuildCreateSuccess(extra string) (GuildCreateSuccess, error) {
	var m GuildCreateSuccess

	if err := m.Deserialize(extra); err != nil {
		return GuildCreateSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildCreateSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildCreateSuccess
}

func (m GuildCreateSuccess) MessageName() string {
	return "GuildCreateSuccess"
}

func (m GuildCreateSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildCreateSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
