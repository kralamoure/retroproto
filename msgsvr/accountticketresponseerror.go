package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountTicketResponseError struct{}

func (m AccountTicketResponseError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountTicketResponseError
}

func (m AccountTicketResponseError) Serialized() (string, error) {
	return "", nil
}

func (m *AccountTicketResponseError) Deserialize(extra string) error {
	return nil
}
