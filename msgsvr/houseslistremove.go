// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type HousesListRemove struct{}

func (m HousesListRemove) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.HousesListRemove
}

func (m HousesListRemove) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesListRemove) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
