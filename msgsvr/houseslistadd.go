// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type HousesListAdd struct{}

func (m HousesListAdd) ProtocolId() retroproto.MsgSvrId {
	return retroproto.HousesListAdd
}

func (m HousesListAdd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesListAdd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
