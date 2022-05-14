// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildJoinError struct{}

func NewGuildJoinError(extra string) (GuildJoinError, error) {
	var m GuildJoinError

	if err := m.Deserialize(extra); err != nil {
		return GuildJoinError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildJoinError) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildJoinError
}

func (m GuildJoinError) MessageName() string {
	return "GuildJoinError"
}

func (m GuildJoinError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildJoinError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
