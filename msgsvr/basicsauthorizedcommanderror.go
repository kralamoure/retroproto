package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedCommandError struct{}

func NewBasicsAuthorizedCommandError(extra string) (BasicsAuthorizedCommandError, error) {
	var m BasicsAuthorizedCommandError

	if err := m.Deserialize(extra); err != nil {
		return BasicsAuthorizedCommandError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsAuthorizedCommandError) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsAuthorizedCommandError
}

func (m BasicsAuthorizedCommandError) MessageName() string {
	return "BasicsAuthorizedCommandError"
}

func (m BasicsAuthorizedCommandError) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsAuthorizedCommandError) Deserialize(extra string) error {
	return nil
}
