// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildCreateError struct{}

func NewGuildCreateError(extra string) (GuildCreateError, error) {
	var m GuildCreateError

	if err := m.Deserialize(extra); err != nil {
		return GuildCreateError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildCreateError) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildCreateError
}

func (m GuildCreateError) MessageName() string {
	return "GuildCreateError"
}

func (m GuildCreateError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildCreateError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
