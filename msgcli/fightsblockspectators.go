// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type FightsBlockSpectators struct{}

func (m FightsBlockSpectators) MessageId() retroproto.MsgCliId {
	return retroproto.FightsBlockSpectators
}

func (m FightsBlockSpectators) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsBlockSpectators) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
