// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeMountStorageRemove struct{}

func (m ExchangeMountStorageRemove) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeMountStorageRemove
}

func (m ExchangeMountStorageRemove) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeMountStorageRemove) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
