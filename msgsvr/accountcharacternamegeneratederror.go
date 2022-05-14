package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type AccountCharacterNameGeneratedError struct {
	Reason int
}

func NewAccountCharacterNameGeneratedError(extra string) (AccountCharacterNameGeneratedError, error) {
	var m AccountCharacterNameGeneratedError

	if err := m.Deserialize(extra); err != nil {
		return AccountCharacterNameGeneratedError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountCharacterNameGeneratedError) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterNameGeneratedError
}

func (m AccountCharacterNameGeneratedError) MessageName() string {
	return "AccountCharacterNameGeneratedError"
}

func (m AccountCharacterNameGeneratedError) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Reason), nil
}

func (m *AccountCharacterNameGeneratedError) Deserialize(extra string) error {
	reason, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Reason = int(reason)

	return nil
}
