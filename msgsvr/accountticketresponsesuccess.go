package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type AccountTicketResponseSuccess struct {
	KeyId int
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
