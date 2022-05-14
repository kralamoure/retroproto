package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type ExchangePutInMountParkFromShed struct {
	MountId int
}

func NewExchangePutInMountParkFromShed(extra string) (ExchangePutInMountParkFromShed, error) {
	var m ExchangePutInMountParkFromShed

	if err := m.Deserialize(extra); err != nil {
		return ExchangePutInMountParkFromShed{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangePutInMountParkFromShed) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangePutInMountParkFromShed
}

func (m ExchangePutInMountParkFromShed) MessageName() string {
	return "ExchangePutInMountParkFromShed"
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
