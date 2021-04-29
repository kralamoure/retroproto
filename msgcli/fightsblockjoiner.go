// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type FightsBlockJoiner struct{}

func (m FightsBlockJoiner) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.FightsBlockJoiner
}

func (m FightsBlockJoiner) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FightsBlockJoiner) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
