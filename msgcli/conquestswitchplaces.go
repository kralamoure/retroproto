// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ConquestSwitchPlaces struct{}

func (m ConquestSwitchPlaces) MessageId() retroproto.MsgCliId {
	return retroproto.ConquestSwitchPlaces
}

func (m ConquestSwitchPlaces) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestSwitchPlaces) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
