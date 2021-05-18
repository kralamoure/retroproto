package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type ExchangePutInMountParkFromShed struct {
	MountId int
}

func (m ExchangePutInMountParkFromShed) ProtocolId() retroproto.MsgCliId {
	return retroproto.ExchangePutInMountParkFromShed
}

func (m ExchangePutInMountParkFromShed) Serialized() (string, error) {
	return fmt.Sprint(m.MountId), nil
}

func (m *ExchangePutInMountParkFromShed) Deserialize(extra string) error {
	mountId, err := strconv.Atoi(extra)
	if err != nil {
		return err
	}
	m.MountId = mountId

	return nil
}
