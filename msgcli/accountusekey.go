package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type AccountUseKey struct {
	Id int
}

func (m AccountUseKey) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountUseKey
}

func (m AccountUseKey) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Id), nil
}

func (m *AccountUseKey) Deserialize(extra string) error {
	id, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	return nil
}
