package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retro/retrotyp"

	"github.com/kralamoure/retroproto"
)

type ExchangeBigStoreType struct {
	ItemType retrotyp.ItemType
}

func NewExchangeBigStoreType(extra string) (ExchangeBigStoreType, error) {
	var m ExchangeBigStoreType

	if err := m.Deserialize(extra); err != nil {
		return ExchangeBigStoreType{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeBigStoreType) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeBigStoreType
}

func (m ExchangeBigStoreType) MessageName() string {
	return "ExchangeBigStoreType"
}

func (m ExchangeBigStoreType) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.ItemType), nil
}

func (m *ExchangeBigStoreType) Deserialize(extra string) error {
	v, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.ItemType = retrotyp.ItemType(v)
	return nil
}
