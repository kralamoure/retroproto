package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountTicketResponseError struct{}

func NewAccountTicketResponseError(extra string) (AccountTicketResponseError, error) {
	var m AccountTicketResponseError

	if err := m.Deserialize(extra); err != nil {
		return AccountTicketResponseError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountTicketResponseError) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountTicketResponseError
}

func (m AccountTicketResponseError) MessageName() string {
	return "AccountTicketResponseError"
}

func (m AccountTicketResponseError) Serialized() (string, error) {
	return "", nil
}

func (m *AccountTicketResponseError) Deserialize(extra string) error {
	return nil
}
