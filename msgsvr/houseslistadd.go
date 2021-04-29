// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type HousesListAdd struct{}

func (m HousesListAdd) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.HousesListAdd
}

func (m HousesListAdd) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesListAdd) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
