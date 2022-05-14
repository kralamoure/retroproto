package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountTicketResponseError struct{}

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
