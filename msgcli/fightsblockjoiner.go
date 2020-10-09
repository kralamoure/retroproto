// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type FightsBlockJoiner struct{}

func (m FightsBlockJoiner) ProtocolId() d1proto.MsgCliId {
	return d1proto.FightsBlockJoiner
}

func (m FightsBlockJoiner) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *FightsBlockJoiner) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
