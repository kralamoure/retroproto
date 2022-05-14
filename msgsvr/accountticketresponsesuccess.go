package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type AccountTicketResponseSuccess struct {
	KeyId int
}

func NewAccountTicketResponseSuccess(extra string) (AccountTicketResponseSuccess, error) {
	var m AccountTicketResponseSuccess

	if err := m.Deserialize(extra); err != nil {
		return AccountTicketResponseSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountTicketResponseSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountTicketResponseSuccess
}

func (m AccountTicketResponseSuccess) MessageName() string {
	return "AccountTicketResponseSuccess"
}

func (m AccountTicketResponseSuccess) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.KeyId), nil
}

func (m *AccountTicketResponseSuccess) Deserialize(extra string) error {
	if len(extra) < 1 {
		return retroproto.ErrInvalidMsg
	}

	keyId, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.KeyId = int(keyId)

	return nil
}
