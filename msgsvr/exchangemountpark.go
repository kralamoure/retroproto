// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeMountPark struct{}

func (m ExchangeMountPark) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeMountPark
}

func (m ExchangeMountPark) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeMountPark) Deserialize(extra string) error {
	return nil
}
