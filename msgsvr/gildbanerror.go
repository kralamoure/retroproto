// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GildBanError struct{}

func NewGildBanError(extra string) (GildBanError, error) {
	var m GildBanError

	if err := m.Deserialize(extra); err != nil {
		return GildBanError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GildBanError) MessageId() retroproto.MsgSvrId {
	return retroproto.GildBanError
}

func (m GildBanError) MessageName() string {
	return "GildBanError"
}

func (m GildBanError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GildBanError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
