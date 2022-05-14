package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/typ"
)

type ExchangeMountStorageAdd struct {
	Data    typ.CommonMountData
	NewBorn bool
}

func NewExchangeMountStorageAdd(extra string) (ExchangeMountStorageAdd, error) {
	var m ExchangeMountStorageAdd

	if err := m.Deserialize(extra); err != nil {
		return ExchangeMountStorageAdd{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeMountStorageAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeMountStorageAdd
}

func (m ExchangeMountStorageAdd) MessageName() string {
	return "ExchangeMountStorageAdd"
}

func (m ExchangeMountStorageAdd) Serialized() (string, error) {
	data, err := m.Data.Serialized()
	if err != nil {
		return "", err
	}

	newBorn := "+"
	if m.NewBorn {
		newBorn = "~"
	}

	return newBorn + data, nil
}

func (m *ExchangeMountStorageAdd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
