// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type HousesBuySuccess struct{}

func (m HousesBuySuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.HousesBuySuccess
}

func (m HousesBuySuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesBuySuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
