// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedCommandSuccess struct{}

func NewBasicsAuthorizedCommandSuccess(extra string) (BasicsAuthorizedCommandSuccess, error) {
	var m BasicsAuthorizedCommandSuccess

	if err := m.Deserialize(extra); err != nil {
		return BasicsAuthorizedCommandSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsAuthorizedCommandSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsAuthorizedCommandSuccess
}

func (m BasicsAuthorizedCommandSuccess) MessageName() string {
	return "BasicsAuthorizedCommandSuccess"
}

func (m BasicsAuthorizedCommandSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedCommandSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
