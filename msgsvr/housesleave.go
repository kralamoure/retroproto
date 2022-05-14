// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type HousesLeave struct{}

func (m HousesLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.HousesLeave
}

func (m HousesLeave) MessageName() string {
	return "HousesLeave"
}

func (m HousesLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
