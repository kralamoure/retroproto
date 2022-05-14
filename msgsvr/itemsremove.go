package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type ItemsRemove struct {
	Id int
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
