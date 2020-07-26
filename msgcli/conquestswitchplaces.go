// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ConquestSwitchPlaces struct{}

func (m ConquestSwitchPlaces) ProtocolId() d1proto.MsgCliId {
	return d1proto.ConquestSwitchPlaces
}

func (m ConquestSwitchPlaces) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ConquestSwitchPlaces) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
