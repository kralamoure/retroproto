// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeMountPods struct{}

func (m ExchangeMountPods) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeMountPods
}

func (m ExchangeMountPods) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeMountPods) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
