package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type ExchangePutInShedFromInventory struct {
	MountId int
}

func (m ExchangePutInShedFromInventory) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangePutInShedFromInventory
}

func (m ExchangePutInShedFromInventory) Serialized() (string, error) {
	return fmt.Sprint(m.MountId), nil
}

func (m *ExchangePutInShedFromInventory) Deserialize(extra string) error {
	mountId, err := strconv.Atoi(extra)
	if err != nil {
		return err
	}
	m.MountId = mountId

	return nil
}
