package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type ExchangePutInShedFromInventory struct {
	MountId int
}

func NewExchangePutInShedFromInventory(extra string) (ExchangePutInShedFromInventory, error) {
	var m ExchangePutInShedFromInventory

	if err := m.Deserialize(extra); err != nil {
		return ExchangePutInShedFromInventory{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangePutInShedFromInventory) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangePutInShedFromInventory
}

func (m ExchangePutInShedFromInventory) MessageName() string {
	return "ExchangePutInShedFromInventory"
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
