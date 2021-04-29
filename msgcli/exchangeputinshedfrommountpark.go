package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type ExchangePutInShedFromMountPark struct {
	MountId int
}

func (m ExchangePutInShedFromMountPark) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangePutInShedFromMountPark
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
