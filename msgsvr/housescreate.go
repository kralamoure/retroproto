// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type HousesCreate struct{}

func (m HousesCreate) MessageId() retroproto.MsgSvrId {
	return retroproto.HousesCreate
}

func (m HousesCreate) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesCreate) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
