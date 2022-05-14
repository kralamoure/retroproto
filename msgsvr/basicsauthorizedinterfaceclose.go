// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedInterfaceClose struct{}

func NewBasicsAuthorizedInterfaceClose(extra string) (BasicsAuthorizedInterfaceClose, error) {
	var m BasicsAuthorizedInterfaceClose

	if err := m.Deserialize(extra); err != nil {
		return BasicsAuthorizedInterfaceClose{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsAuthorizedInterfaceClose) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsAuthorizedInterfaceClose
}

func (m BasicsAuthorizedInterfaceClose) MessageName() string {
	return "BasicsAuthorizedInterfaceClose"
}

func (m BasicsAuthorizedInterfaceClose) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedInterfaceClose) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
