package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountCharacterAddSuccess struct{}

func NewAccountCharacterAddSuccess(extra string) (AccountCharacterAddSuccess, error) {
	var m AccountCharacterAddSuccess

	if err := m.Deserialize(extra); err != nil {
		return AccountCharacterAddSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountCharacterAddSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterAddSuccess
}

func (m AccountCharacterAddSuccess) MessageName() string {
	return "AccountCharacterAddSuccess"
}

func (m AccountCharacterAddSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *AccountCharacterAddSuccess) Deserialize(extra string) error {
	return nil
}
