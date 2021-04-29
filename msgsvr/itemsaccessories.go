package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1encoding"
	"github.com/kralamoure/d1encoding/typ"
)

type ItemsAccessories struct {
	Id          int
	Accessories typ.CommonAccessories
}

func (m ItemsAccessories) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ItemsAccessories
}

func (m ItemsAccessories) Serialized() (string, error) {
	accessories, err := m.Accessories.Serialized()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d|%s", m.Id, accessories), nil
}

func (m *ItemsAccessories) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 2 {
		return d1encoding.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	err = m.Accessories.Deserialize(sli[1])
	if err != nil {
		return err
	}

	return nil
}
