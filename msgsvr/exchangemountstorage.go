// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeMountStorage struct{}

func (m ExchangeMountStorage) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeMountStorage
}

func (m ExchangeMountStorage) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeMountStorage) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
