package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountTicketResponseError struct{}

func (m AccountTicketResponseError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountTicketResponseError
}

func (m AccountTicketResponseError) Serialized() (string, error) {
	return "", nil
}

func (m *AccountTicketResponseError) Deserialize(extra string) error {
	return nil
}
