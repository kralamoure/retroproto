// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsWhoIsSuccess struct{}

func NewBasicsWhoIsSuccess(extra string) (BasicsWhoIsSuccess, error) {
	var m BasicsWhoIsSuccess

	if err := m.Deserialize(extra); err != nil {
		return BasicsWhoIsSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsWhoIsSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsWhoIsSuccess
}

func (m BasicsWhoIsSuccess) MessageName() string {
	return "BasicsWhoIsSuccess"
}

func (m BasicsWhoIsSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsWhoIsSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
