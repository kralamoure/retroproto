// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeMountPods struct{}

func (m ExchangeMountPods) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeMountPods
}

func (m ExchangeMountPods) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeMountPods) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
