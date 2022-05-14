package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type ExchangePutInShedFromMountPark struct {
	MountId int
}

func NewExchangePutInShedFromMountPark(extra string) (ExchangePutInShedFromMountPark, error) {
	var m ExchangePutInShedFromMountPark

	if err := m.Deserialize(extra); err != nil {
		return ExchangePutInShedFromMountPark{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangePutInShedFromMountPark) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangePutInShedFromMountPark
}

func (m ExchangePutInShedFromMountPark) MessageName() string {
	return "ExchangePutInShedFromMountPark"
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
