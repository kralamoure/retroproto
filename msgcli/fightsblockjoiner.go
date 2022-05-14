// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type FightsBlockJoiner struct{}

func (m FightsBlockJoiner) MessageId() retroproto.MsgCliId {
	return retroproto.FightsBlockJoiner
}

func (m FightsBlockJoiner) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsBlockJoiner) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
