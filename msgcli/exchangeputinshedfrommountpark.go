package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type ExchangePutInShedFromMountPark struct {
	MountId int
}

func (m ExchangePutInShedFromMountPark) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangePutInShedFromMountPark
}

func (m ExchangePutInShedFromMountPark) Serialized() (string, error) {
	return fmt.Sprint(m.MountId), nil
}

func (m *ExchangePutInShedFromMountPark) Deserialize(extra string) error {
	mountId, err := strconv.Atoi(extra)
	if err != nil {
		return err
	}
	m.MountId = mountId

	return nil
}
