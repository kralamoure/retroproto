// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type HousesBuyError struct{}

func (m HousesBuyError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.HousesBuyError
}

func (m HousesBuyError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesBuyError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
