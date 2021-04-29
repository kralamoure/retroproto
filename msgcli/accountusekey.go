package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type AccountUseKey struct {
	Id int
}

func (m AccountUseKey) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountUseKey
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
