package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type ItemsDestroy struct {
	Id       int
	Quantity int
}

func (m ItemsDestroy) ProtocolId() d1proto.MsgCliId {
	return d1proto.ItemsDestroy
}

func (m ItemsDestroy) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.Id, m.Quantity), nil
}

func (m *ItemsDestroy) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) != 2 {
		return d1proto.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	quantity, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.Quantity = int(quantity)

	return nil
}
