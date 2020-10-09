// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type FightsGetList struct{}

func (m FightsGetList) ProtocolId() d1proto.MsgCliId {
	return d1proto.FightsGetList
}

func (m FightsGetList) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *FightsGetList) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
