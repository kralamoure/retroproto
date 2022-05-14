// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedInterfaceOpen struct{}

func NewBasicsAuthorizedInterfaceOpen(extra string) (BasicsAuthorizedInterfaceOpen, error) {
	var m BasicsAuthorizedInterfaceOpen

	if err := m.Deserialize(extra); err != nil {
		return BasicsAuthorizedInterfaceOpen{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsAuthorizedInterfaceOpen) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsAuthorizedInterfaceOpen
}

func (m BasicsAuthorizedInterfaceOpen) MessageName() string {
	return "BasicsAuthorizedInterfaceOpen"
}

func (m BasicsAuthorizedInterfaceOpen) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedInterfaceOpen) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
