package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountTicketResponseError struct{}

func (m AccountTicketResponseError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AccountTicketResponseError
}

func (m AccountTicketResponseError) Serialized() (string, error) {
	return "", nil
}

func (m *AccountTicketResponseError) Deserialize(extra string) error {
	return nil
}
