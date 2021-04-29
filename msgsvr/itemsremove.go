package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type ItemsRemove struct {
	Id int
}

func (m ItemsRemove) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsRemove
}

func (m ItemsRemove) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Id), nil
}

func (m *ItemsRemove) Deserialize(extra string) error {
	if extra == "" {
		return d1proto.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	return nil
}
