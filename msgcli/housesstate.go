// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type HousesState struct{}

func (m HousesState) MessageId() retroproto.MsgCliId {
	return retroproto.HousesState
}

func (m HousesState) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesState) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
