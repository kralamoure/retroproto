// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type KeyKeyError struct{}

func NewKeyKeyError(extra string) (KeyKeyError, error) {
	var m KeyKeyError

	if err := m.Deserialize(extra); err != nil {
		return KeyKeyError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m KeyKeyError) MessageId() retroproto.MsgSvrId {
	return retroproto.KeyKeyError
}

func (m KeyKeyError) MessageName() string {
	return "KeyKeyError"
}

func (m KeyKeyError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *KeyKeyError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
