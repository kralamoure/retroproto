package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type AccountCommunity struct {
	Id int
}

func NewAccountCommunity(extra string) (AccountCommunity, error) {
	var m AccountCommunity

	if err := m.Deserialize(extra); err != nil {
		return AccountCommunity{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountCommunity) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCommunity
}

func (m AccountCommunity) MessageName() string {
	return "AccountCommunity"
}

func (m AccountCommunity) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Id), nil
}

func (m *AccountCommunity) Deserialize(extra string) error {
	if len(extra) < 1 {
		return retroproto.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	return nil
}
