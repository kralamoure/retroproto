package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type AccountUseKey struct {
	Id int
}

func NewAccountUseKey(extra string) (AccountUseKey, error) {
	var m AccountUseKey

	if err := m.Deserialize(extra); err != nil {
		return AccountUseKey{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountUseKey) MessageId() retroproto.MsgCliId {
	return retroproto.AccountUseKey
}

func (m AccountUseKey) MessageName() string {
	return "AccountUseKey"
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
