package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type ExchangePutInInventoryFromShed struct {
	MountId int
}

func (m ExchangePutInInventoryFromShed) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangePutInInventoryFromShed
}

func (m ExchangePutInInventoryFromShed) Serialized() (string, error) {
	return fmt.Sprint(m.MountId), nil
}

func (m *ExchangePutInInventoryFromShed) Deserialize(extra string) error {
	mountId, err := strconv.Atoi(extra)
	if err != nil {
		return err
	}
	m.MountId = mountId

	return nil
}
