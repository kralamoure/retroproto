// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsFileCheck struct{}

func NewBasicsFileCheck(extra string) (BasicsFileCheck, error) {
	var m BasicsFileCheck

	if err := m.Deserialize(extra); err != nil {
		return BasicsFileCheck{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsFileCheck) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsFileCheck
}

func (m BasicsFileCheck) MessageName() string {
	return "BasicsFileCheck"
}

func (m BasicsFileCheck) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsFileCheck) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
