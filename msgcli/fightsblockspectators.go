// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type FightsBlockSpectators struct{}

func (m FightsBlockSpectators) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.FightsBlockSpectators
}

func (m FightsBlockSpectators) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FightsBlockSpectators) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
