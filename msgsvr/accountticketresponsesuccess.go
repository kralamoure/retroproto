package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type AccountTicketResponseSuccess struct {
	KeyId int
}

func (m AccountTicketResponseSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountTicketResponseSuccess
}

func (m AccountTicketResponseSuccess) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.KeyId), nil
}

func (m *AccountTicketResponseSuccess) Deserialize(extra string) error {
	if len(extra) < 1 {
		return d1encoding.ErrInvalidMsg
	}

	keyId, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.KeyId = int(keyId)

	return nil
}
