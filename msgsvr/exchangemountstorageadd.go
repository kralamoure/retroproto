package msgsvr

import (
	"github.com/kralamoure/d1proto"
	"github.com/kralamoure/d1proto/typ"
)

type ExchangeMountStorageAdd struct {
	Data    typ.CommonMountData
	NewBorn bool
}

func (m ExchangeMountStorageAdd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeMountStorageAdd
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
	return d1proto.ErrNotImplemented
}
