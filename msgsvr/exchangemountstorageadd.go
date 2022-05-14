package msgsvr

import (
	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/typ"
)

type ExchangeMountStorageAdd struct {
	Data    typ.CommonMountData
	NewBorn bool
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
