// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsWhoIsError struct{}

func NewBasicsWhoIsError(extra string) (BasicsWhoIsError, error) {
	var m BasicsWhoIsError

	if err := m.Deserialize(extra); err != nil {
		return BasicsWhoIsError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsWhoIsError) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsWhoIsError
}

func (m BasicsWhoIsError) MessageName() string {
	return "BasicsWhoIsError"
}

func (m BasicsWhoIsError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsWhoIsError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
