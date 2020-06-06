// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type FightsBlockSpectators struct{}

func (m FightsBlockSpectators) ProtocolId() d1proto.MsgCliId {
	return d1proto.FightsBlockSpectators
}

func (m FightsBlockSpectators) Serialized() (string, error) {
	return "", nil
}

func (m *FightsBlockSpectators) Deserialize(extra string) error {
	return nil
}
