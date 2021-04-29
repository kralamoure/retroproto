package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type ExchangePutInShedFromInventory struct {
	MountId int
}

func (m ExchangePutInShedFromInventory) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangePutInShedFromInventory
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
