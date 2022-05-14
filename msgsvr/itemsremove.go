package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type ItemsRemove struct {
	Id int
}

func NewItemsRemove(extra string) (ItemsRemove, error) {
	var m ItemsRemove

	if err := m.Deserialize(extra); err != nil {
		return ItemsRemove{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ItemsRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.ItemsRemove
}

func (m ItemsRemove) MessageName() string {
	return "ItemsRemove"
}

func (m ItemsRemove) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Id), nil
}

func (m *ItemsRemove) Deserialize(extra string) error {
	if extra == "" {
		return retroproto.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	return nil
}
