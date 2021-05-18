package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type AccountUseKey struct {
	Id int
}

func (m AccountUseKey) ProtocolId() retroproto.MsgCliId {
	return retroproto.AccountUseKey
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
