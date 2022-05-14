// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountCharacterDeleteSuccess struct{}

func NewAccountCharacterDeleteSuccess(extra string) (AccountCharacterDeleteSuccess, error) {
	var m AccountCharacterDeleteSuccess

	if err := m.Deserialize(extra); err != nil {
		return AccountCharacterDeleteSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountCharacterDeleteSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterDeleteSuccess
}

func (m AccountCharacterDeleteSuccess) MessageName() string {
	return "AccountCharacterDeleteSuccess"
}

func (m AccountCharacterDeleteSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountCharacterDeleteSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
