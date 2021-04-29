// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type HousesSellSuccess struct{}

func (m HousesSellSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.HousesSellSuccess
}

func (m HousesSellSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesSellSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
