package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type ExchangePutInInventoryFromShed struct {
	MountId int
}

func NewExchangePutInInventoryFromShed(extra string) (ExchangePutInInventoryFromShed, error) {
	var m ExchangePutInInventoryFromShed

	if err := m.Deserialize(extra); err != nil {
		return ExchangePutInInventoryFromShed{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangePutInInventoryFromShed) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangePutInInventoryFromShed
}

func (m ExchangePutInInventoryFromShed) MessageName() string {
	return "ExchangePutInInventoryFromShed"
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
