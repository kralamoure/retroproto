// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestSwitchPlaces struct{}

func (m ConquestSwitchPlaces) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ConquestSwitchPlaces
}

func (m ConquestSwitchPlaces) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestSwitchPlaces) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
