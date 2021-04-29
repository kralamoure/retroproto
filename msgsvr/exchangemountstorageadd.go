package msgsvr

import (
	"github.com/kralamoure/d1encoding"
	"github.com/kralamoure/d1encoding/typ"
)

type ExchangeMountStorageAdd struct {
	Data    typ.CommonMountData
	NewBorn bool
}

func (m ExchangeMountStorageAdd) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeMountStorageAdd
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
	return d1encoding.ErrNotImplemented
}
