// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeMountPods struct{}

func (m ExchangeMountPods) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ExchangeMountPods
}

func (m ExchangeMountPods) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeMountPods) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
